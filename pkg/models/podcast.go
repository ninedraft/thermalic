package models

import (
	"time"
)

type Podcast struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EpisodesNum uint64    `json:"series_num"`
	FeedURL     string    `json:"feed_url"`
	LastUpdated time.Time `json:"last_updated"`
	Downloaded  uint64    `json:"downloaded"`
	Listened    uint64    `json:"listened"`
	Episodes    []Episode `json:"episodes"`
}
