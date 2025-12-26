package model

import "gorm.io/gorm"

// Channel represents an API channel
type Channel struct {
	gorm.Model
	Type    int    `json:"type"`
	Key     string `json:"key"`
	BaseURL string `json:"base_url"`
	Name    string `json:"name"`
	// Additional fields for local channel
	IsLocal    bool   `json:"is_local"`
	Command    string `json:"command"` // For local binary runner
	Parameters string `json:"parameters"`
}
