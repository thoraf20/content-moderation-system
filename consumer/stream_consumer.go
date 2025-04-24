package consumer

import (
	"context"
	"log"
	"time"

	"content-analysis/config"
	"content-analysis/moderation"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func StartStreamConsumer() {
	client := redis.NewClient(&redis.Options{
		Addr: config.Get("BROKER_URL"),
	})

	streamName := config.Get("TOPIC_NAME")
	lastID := "0"

	for {
		streams, err := client.XRead(ctx, &redis.XReadArgs{
			Streams: []string{streamName, lastID},
			Block:   5 * time.Second,
			Count:   1,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Printf("Read error: %v", err)
			continue
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				data := message.Values["event"].(string)

				log.Printf("Moderating event: %s", data)
				moderation.ProcessEvent(data)

				lastID = message.ID
			}
		}
	}
}
