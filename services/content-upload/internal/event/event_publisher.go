package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()

type ModerationEvent struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Type     string `json:"type"`
	Content  string `json:"content,omitempty"`
}

func PublishModerationEvent(filename, path, fileType string, rawText ...string) {
	broker := viper.GetString("BROKER_URL")
	stream := viper.GetString("TOPIC_NAME")

	client := redis.NewClient(&redis.Options{
		Addr: broker,
	})

	defer client.Close()

	event := ModerationEvent{
		Filename: filename,
		Path:     path,
		Type:     fileType,
	}

	if len(rawText) > 0 {
		event.Content = rawText[0] // Optional: Add `Content` field to ModerationEvent struct
	}

	data, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Failed to marshal event: %v", err)
		return
	}

	// Add to stream
	err = client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream, // 
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
