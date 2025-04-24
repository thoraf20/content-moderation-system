package consumer

import (
	"context"
	"encoding/json"
	"log"
	"content-analysis/moderation"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

type ModerationEvent struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Type     string `json:"type"` // "text", "image", or "video"
	Content  string `json:"content"` // optional, used for text moderation
}

func StartStreamConsumer() {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("BROKER_URL"),
	})

	stream := os.Getenv("TOPIC_NAME")

	log.Println("Listening to Redis Stream:", stream)

	for {
		streams, err := client.XRead(ctx, &redis.XReadArgs{
			Streams: []string{stream, "0"},
			Count:   1,
			Block:   0,
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

				log.Printf("Processing event: %+v\n", event)

				switch event.Type {
				case "text":
					isClean, reason := moderation.ModerateText(event.Content)
					if !isClean {
						log.Printf("Text moderation failed: %s", reason)
					} else {
						log.Printf("Text passed moderation")
					}
				default:
					log.Printf("Unsupported content type: %s", event.Type)
				}
			}
		}
	}
}
