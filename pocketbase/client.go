package pocketbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	Token          string
	BaseURL        string
	CollectionName string
	HTTPClient     *http.Client
}

func NewClient(baseURL string, collectionName string) *Client {
	return &Client{
		BaseURL:        baseURL,
		CollectionName: collectionName,
		HTTPClient:     &http.Client{},
	}
}

func (c *Client) Login(ctx context.Context, identity, password string) error {
	url := fmt.Sprintf("%s/api/collections/_superusers/auth-with-password", c.BaseURL)

	requestBody := map[string]string{
		"identity": identity,
		"password": password,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var loginResponse struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		return err
	}
	c.Token = loginResponse.Token
	return nil
}

func (c *Client) GetAllMusicData(ctx context.Context) ([]MaiMaiMusicData, error) {
	allMusicData := []MaiMaiMusicData{}
	page := 1
	perPage := 100
	skipTotal := false

	musicData, totalPages, err := c.getMusicData(ctx, page, perPage, skipTotal)
	if err != nil {
		return nil, err
	}
	allMusicData = append(allMusicData, musicData...)
	page++
	skipTotal = true
	for page <= totalPages {
		musicData, totalPages, err = c.getMusicData(ctx, page, perPage, skipTotal)
		if err != nil {
			return nil, err
		}
		allMusicData = append(allMusicData, musicData...)
		page++
	}

	return allMusicData, nil
}

func (c *Client) getMusicData(ctx context.Context, page int, perPage int, skipTotal bool) ([]MaiMaiMusicData, int, error) {
	requestURL := fmt.Sprintf("%s/api/collections/%s/records", c.BaseURL, c.CollectionName)

	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("perPage", strconv.Itoa(perPage))
	params.Add("skipTotal", strconv.FormatBool(skipTotal))
	requestURL += "?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	type MusicDataResponse struct {
		TotalPages int               `json:"totalPages"`
		Items      []MaiMaiMusicData `json:"items"`
	}

	var musicDataResponse MusicDataResponse
	if err := json.Unmarshal(body, &musicDataResponse); err != nil {
		return nil, 0, err
	}

	return musicDataResponse.Items, musicDataResponse.TotalPages, nil
}

func (c *Client) BatchUpsertMusicData(ctx context.Context, musicDataListToCreate []MaiMaiMusicData, musicDataListToUpdate []MaiMaiMusicData) error {
	requestURL := fmt.Sprintf("%s/api/batch", c.BaseURL)
	endpoint := fmt.Sprintf("/api/collections/%s/records", c.CollectionName)

	batchRequest := BatchRequest{
		Requests: []BatchRequestItem{},
	}
	for _, musicData := range musicDataListToCreate {
		musicDataToCreate := musicData.ToCreate()
		jsonBody, err := json.Marshal(musicDataToCreate)
		if err != nil {
			return err
		}
		request := BatchRequestItem{
			Method: "POST",
			Url:    endpoint,
			Body:   jsonBody,
		}
		batchRequest.Requests = append(batchRequest.Requests, request)
	}
	for _, musicData := range musicDataListToUpdate {
		musicDataToUpdate := musicData.ToUpdate()
		jsonBody, err := json.Marshal(musicDataToUpdate)
		if err != nil {
			return err
		}
		request := BatchRequestItem{
			Method: "PUT",
			Url:    fmt.Sprintf("%s/%s", endpoint, musicData.ID),
			Body:   jsonBody,
		}
		batchRequest.Requests = append(batchRequest.Requests, request)
	}

	jsonBody, err := json.Marshal(batchRequest)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var batchResponse BatchResponse
	if err := json.Unmarshal(body, &batchResponse); err != nil {
		return err
	}

	for _, response := range batchResponse.Responses {
		if response.Status != http.StatusOK {
			log.Printf("failed to upsert music data: %d, Title: %s, ID: %s\n", response.Status, response.Body.Title, response.Body.ID)
		}
	}

	return nil
}
