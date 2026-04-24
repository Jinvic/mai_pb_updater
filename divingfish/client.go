package divingfish

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mai_pb_updater/config"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Etag       string
}

var defaultBaseURL = "https://www.diving-fish.com/api/maimaidxprober"

func NewClient(baseURL string, etag string) *Client {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
		Etag:       etag,
	}
}

// 获取所有歌曲数据
func (c *Client) GetMaiMaiMusicData(ctx context.Context) ([]MaiMaiMusicData, error) {
	url := fmt.Sprintf("%s/music_data", c.BaseURL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("If-None-Match", c.Etag)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotModified {
			// 没有修改，返回空列表
			return []MaiMaiMusicData{}, nil
		}
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	c.Etag = resp.Header.Get("etag")
	go func() {
		if err := c.presistEtag(); err != nil {
			log.Printf("failed to persist etag: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var musicDataList []MaiMaiMusicData
	if err := json.Unmarshal(body, &musicDataList); err != nil {
		return nil, err
	}
	return musicDataList, nil
}

func (c *Client) presistEtag() error {
	return config.WriteConfig("divingfish.etag", c.Etag)
}
