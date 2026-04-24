package pocketbase

import (
	"context"
	"mai_pb_updater/config"
	"testing"
)

func TestGetAllMusicData(t *testing.T) {
	cfg, err := config.InitConfig("../")
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}

	client := NewClient(cfg.Pocketbase.BaseURL, cfg.Pocketbase.CollectionName, cfg.Pocketbase.MaxBatchSize)
	ctx := context.Background()
	err = client.Login(ctx, cfg.Pocketbase.Identity, cfg.Pocketbase.Password)
	if err != nil {
		t.Fatalf("failed to login: %v", err)
	}
	musicDataList, err := client.GetAllMusicData(ctx)
	if err != nil {
		t.Fatalf("failed to get all music data: %v", err)
	}
	t.Logf("music data list length: %d", len(musicDataList))
}

func TestBatchUpsertMusicData(t *testing.T) {
	cfg, err := config.InitConfig("../")
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}

	client := NewClient(cfg.Pocketbase.BaseURL, cfg.Pocketbase.CollectionName, cfg.Pocketbase.MaxBatchSize)
	ctx := context.Background()
	err = client.Login(ctx, cfg.Pocketbase.Identity, cfg.Pocketbase.Password)
	if err != nil {
		t.Fatalf("failed to login: %v", err)
	}

	musicDataListToCreate := []MaiMaiMusicData{
		{
			ID:     "1",
			Title:  "test",
			Type:   "test",
			DS:     []float64{1, 2, 3},
			Level:  []string{"1", "2", "3"},
			Cids:   []int{1, 2, 3},
			Artist: "test",
		},
		{
			ID:     "2",
			Title:  "test2",
			Type:   "test2",
			DS:     []float64{4, 5, 6},
			Level:  []string{"4", "5", "6"},
			Cids:   []int{4, 5, 6},
			Artist: "test2",
		},
	}

	// musicDataListToUpdate := []MaiMaiMusicData{
	// 	{
	// 		ID: "1",
	// 		Title: "test",
	// 		Type: "test",
	// 		DS: []float64{1, 2, 3},
	// 		Level: []string{"1", "2", "3"},
	// 	},
	// }

	err = client.BatchUpsertMusicData(ctx, musicDataListToCreate, nil)
	if err != nil {
		t.Fatalf("failed to batch upsert music data: %v", err)
	}
	t.Logf("batch upsert music data success")
}
