package pocketbase

import (
	"encoding/json"
	"mai_pb_updater/divingfish"
)

type MaiMaiMusicData struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	DS     DS     `json:"ds"`
	Level  Level  `json:"level"`
	Cids   Cids   `json:"cids"`
	Artist string `json:"artist"`
	Genre  string `json:"genre"`
	BPM    int    `json:"bpm"`
	From   string `json:"from"`
	IsNew  bool   `json:"is_new"`
}

type MaiMaiMusicDataToCreate MaiMaiMusicData

// 只更新部分字段
type MaiMaiMusicDataToUpdate struct {
	DS    DS    `json:"ds"`
	Level Level `json:"level"`
	IsNew bool  `json:"is_new"`
}

func (m *MaiMaiMusicData) FromDivingFish(musicData divingfish.MaiMaiMusicData) {
	m.ID = musicData.ID
	m.Title = musicData.Title
	m.Type = musicData.Type
	m.DS = mapByDifficulty(musicData.DS, musicData.BasicInfo.Genre)
	m.Level = mapByDifficulty(musicData.Level, musicData.BasicInfo.Genre)
	m.Cids = mapByDifficulty(musicData.Cids, musicData.BasicInfo.Genre)
	m.Artist = musicData.BasicInfo.Artist
	m.Genre = musicData.BasicInfo.Genre
	m.BPM = musicData.BasicInfo.BPM
	m.From = musicData.BasicInfo.From
	m.IsNew = musicData.BasicInfo.IsNew
}

func BatchFromDivingFish(musicDataList []divingfish.MaiMaiMusicData) []MaiMaiMusicData {
	newList := make([]MaiMaiMusicData, len(musicDataList))
	for i, musicData := range musicDataList {
		newList[i].FromDivingFish(musicData)
	}
	return newList
}

func (m *MaiMaiMusicData) ToCreate() MaiMaiMusicDataToCreate {
	return MaiMaiMusicDataToCreate(*m)
}

func (m *MaiMaiMusicData) ToUpdate() MaiMaiMusicDataToUpdate {
	return MaiMaiMusicDataToUpdate{
		DS:    m.DS,
		Level: m.Level,
		IsNew: m.IsNew,
	}
}

// --------------------------------------------------------
// Batch Request and Response
// --------------------------------------------------------

type BatchRequest struct {
	Requests []BatchRequestItem `json:"requests"`
}

type BatchRequestItem struct {
	Method string          `json:"method"`
	Url    string          `json:"url"`
	Body   json.RawMessage `json:"body"`
}

type BatchResponse []BatchResponseItem

type BatchResponseItem struct {
	Status int             `json:"status"`
	Body   MaiMaiMusicData `json:"body"`
}
