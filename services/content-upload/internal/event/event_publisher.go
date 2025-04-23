package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type ModerationEvent struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Type     string `json:"type"` // e.g. "image", "video"
}

func PublishModerationEvent(filename, path string) {
	broker := viper.GetString("BROKER_URL")
	topic := viper.GetString("TOPIC_NAME")

	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	event := ModerationEvent{
		Filename: filename,
		Path:     path,
		Type:     "image", // For now — you can add detection later
	}

	data, _ := json.Marshal(event)

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(filename),
			Value: data,
		},
	)

	if err != nil {
		log.Printf("Failed to publish message: %v", err)
	} else {
		log.Printf("Published event for %s", filename)
	}

	w.Close()
}
