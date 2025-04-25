package event

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type ModerationEvent struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Type     string `json:"type"`
}

func PublishModerationEvent(filename, path string, fileType string) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("BROKER_URL"), // "localhost:6379", // Default Redis port
	})

	event := ModerationEvent{
		Filename: filename,
		Path:     path,
		Type:     fileType,
	}

	data, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Failed to marshal event: %v", err)
		return
	}

	// Add to stream
	err = client.XAdd(ctx, &redis.XAddArgs{
		Stream: os.Getenv("TOPIC_URL"), // 
		Values: map[string]interface{}{
			"event": data,
		},
	}).Err()

	if err != nil {
		log.Printf("Failed to publish event: %v", err)
	} else {
		log.Printf("Published event to stream for %s", filename)
	}
}
