package models

import (
	"time"
)

type Episode struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	ID          uint64    `json:"id"`
	Description string    `json:"description"`
	AudioURL    string    `json:"audio_url"`
}
