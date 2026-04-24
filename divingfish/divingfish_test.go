package divingfish

import (
	"context"
	"mai_pb_updater/config"
	"testing"
)

func TestGetMaiMaiMusicData(t *testing.T) {
	cfg, err := config.InitConfig("../")
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}

	// test api
	client := NewClient(cfg.DivingFish.BaseURL, cfg.DivingFish.Etag)
	client.Etag = "" // reset etag
	ctx := context.Background()
	musicDataList, err := client.GetMaiMaiMusicData(ctx)
	if err != nil {
		t.Fatalf("failed to get mai mai music data: %v", err)
	}

	t.Logf("first request: music data list length: %d", len(musicDataList))
	if len(musicDataList) == 0 {
		t.Fatalf("music data list is empty")
	}

	// test etag
	musicDataList, err = client.GetMaiMaiMusicData(ctx)
	if err != nil {
		t.Fatalf("failed to get mai mai music data: %v", err)
	}
	t.Logf("second request: music data list length: %d", len(musicDataList))
	if len(musicDataList) != 0 {
		t.Fatalf("etag is not working")
	}
}
