package main

import (
	"context"
	"flag"
	"log"
	"mai_pb_updater/config"
	"mai_pb_updater/divingfish"
	"mai_pb_updater/pocketbase"
	"os"
)

var configPath string

// 配置优先级: 环境变量 > 命令行参数 > 默认值
func main() {
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flag.StringVar(&configPath, "config", "./config.yml", "config file path")
		flag.Parse()
	}

	cfg, err := config.InitConfig(configPath)
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	updateMusicData(cfg)
}

func updateMusicData(cfg *config.Config) {
	divingfishClient := divingfish.NewClient(cfg.DivingFish.BaseURL, cfg.DivingFish.Etag)
	pocketbaseClient := pocketbase.NewClient(cfg.Pocketbase.BaseURL, cfg.Pocketbase.CollectionName, cfg.Pocketbase.MaxBatchSize)
	ctx := context.Background()
	if err := pocketbaseClient.Login(ctx, cfg.Pocketbase.Identity, cfg.Pocketbase.Password); err != nil {
		log.Fatalf("failed to login to pocketbase: %v", err)
	}

	rawMusicDataList, err := divingfishClient.GetMaiMaiMusicData(ctx)
	if err != nil {
		log.Fatalf("failed to get mai mai music data: %v", err)
	}
	if len(rawMusicDataList) == 0 {
		log.Printf("no music data to update")
		return
	}
	musicDataList := pocketbase.BatchFromDivingFish(rawMusicDataList)

	existingMusicDataList, err := pocketbaseClient.GetAllMusicData(ctx)
	if err != nil {
		log.Fatalf("failed to get all music data: %v", err)
	}
	existingMusicDataMap := make(map[string]struct{})
	for _, musicData := range existingMusicDataList {
		existingMusicDataMap[musicData.ID] = struct{}{}
	}

	musicDataListToCreate := []pocketbase.MaiMaiMusicData{}
	musicDataListToUpdate := []pocketbase.MaiMaiMusicData{}
	for _, musicData := range musicDataList {
		if _, exists := existingMusicDataMap[musicData.ID]; exists {
			musicDataListToUpdate = append(musicDataListToUpdate, musicData)
		} else {
			musicDataListToCreate = append(musicDataListToCreate, musicData)
		}
	}

	if err := pocketbaseClient.BatchUpsertMusicData(ctx, musicDataListToCreate, musicDataListToUpdate); err != nil {
		log.Fatalf("failed to batch upsert music data: %v", err)
	}

	log.Printf("created %d music data", len(musicDataListToCreate))
	log.Printf("updated %d music data", len(musicDataListToUpdate))
}
