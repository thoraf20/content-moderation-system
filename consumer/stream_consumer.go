package consumer

import (
	"content-analysis/config"
	"content-analysis/moderation"
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type ModerationEvent struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Type     string `json:"type"` // "text", "image", or "video"
	Content  string `json:"content"` // optional, used for text moderation
}

func StartStreamConsumer() {
	config.LoadConfig()

	client := redis.NewClient(&redis.Options{
		Addr: config.Get("BROKER_URL"),
	})

	defer client.Close()

	lastID := "0"

	for {
		streams, err := client.XRead(ctx, &redis.XReadArgs{
			Streams: []string{config.Get("TOPIC_NAME"), lastID},
			Count:   0,
			Block:   10,
		}).Result()

		if err != nil {
			log.Printf("Error reading stream: %v", err)
			continue
		}

		for _, stream := range streams {
			for _, msg := range stream.Messages {
				eventJSON := msg.Values["event"].(string)

				var event ModerationEvent
				if err := json.Unmarshal([]byte(eventJSON), &event); 
				err != nil {
					log.Printf("Error unmarshaling event: %v", err)
					continue
				}

				switch event.Type {
				case "text":
					go moderation.ModerateText(event.Content, event.Filename)
				case "image":
					go moderation.ModerateText(event.Path, event.Filename)
				case "video":
					go moderation.ModerateText(event.Path, event.Filename)
				default:
					log.Printf("Unknown content type: %s", event.Type)
				}

				lastID = msg.ID
			}
		}
	}
}
