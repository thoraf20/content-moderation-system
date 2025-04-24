package moderation

import (
	"encoding/json"
	"log"
	"strings"
)

type ModerationEvent struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Type     string `json:"type"`
}

func ProcessEvent(data string) {
	var event ModerationEvent
	if err := json.Unmarshal([]byte(data), &event); 
	err != nil {
		log.Printf("Failed to parse event: %v", err)
		return
	}

	if strings.Contains(strings.ToLower(event.Filename), "banned") {
		log.Printf("❌ Rejected: %s (banned keyword)", event.Filename)
	} else {
		log.Printf("✅ Approved: %s", event.Filename)
	}
}
