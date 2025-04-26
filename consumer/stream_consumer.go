package consumer

import (
	"content-analysis/config"
	"content-analysis/moderation"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

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

	// Create the consumer group if it doesn't exist
	err := client.XGroupCreateMkStream(ctx, config.Get("TOPIC_NAME"), "moderation_group", "$").Err()
	defer client.Close()

	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Fatalf("Could not create consumer group: %v", err)
	}

	consumerName := fmt.Sprintf("consumer-%d", time.Now().UnixNano())

	for {
		claimPendingMessages(client, consumerName)
		// streams, err := client.XRead(ctx, &redis.XReadArgs{
		streams, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "moderation_group",
			Consumer: consumerName,
			Streams: []string{config.Get("TOPIC_NAME"), ">"},
			Count:   5,
			Block:   0,
		}).Result()

		if err != nil {
			log.Printf("Error reading stream: %v", err)
			continue
		}

		for _, stream := range streams {
			for _, msg := range stream.Messages {
				processAndAckMessage(client, msg)
			}
		}
	}
}

func processAndAckMessage(client *redis.Client, msg redis.XMessage) {
	eventJSON, ok := msg.Values["event"].(string)
	if !ok {
		log.Println("Invalid event format")
		return
	}

	var event ModerationEvent
	if err := json.Unmarshal([]byte(eventJSON), &event); 
	err != nil {
		log.Printf("Error unmarshaling event: %v", err)
		return
	}

	engine, err := moderation.GetModerationEngine(event.Type)
	if err != nil {
		log.Printf("Error getting moderation engine: %v", err)
		return
	}

	var content string
	if event.Type == "text" {
		content = event.Content
	} else {
		content = event.Path
	}

	err = engine.Moderate(content, event.Filename)
	if err == nil {
		err := client.XAck(ctx, config.Get("TOPIC_NAME"), "moderation_group", msg.ID).Err()
		if err != nil {
			log.Printf("Failed to ACK message: %v", err)
		} else {
			log.Printf("ACKed message ID: %s", msg.ID)
		}
	} else {
		log.Printf("Moderation failed for message ID %s: %v", msg.ID, err)
	}
}

func claimPendingMessages(client *redis.Client, consumerName string) {
	pending, err := client.XPendingExt(ctx, &redis.XPendingExtArgs{
    Stream: config.Get("TOPIC_NAME"),
    Group:  "moderation_group",
    Start:  "-",
    End:    "+",
    Count:  5,
	}).Result()

	if err != nil && err != redis.Nil {
		log.Printf("Error checking pending: %v", err)
		return
	}

	for _, pend := range pending {
		// Claim message
		claimedMsgs, err := client.XClaim(ctx, &redis.XClaimArgs{
			Stream:   config.Get("TOPIC_NAME"),
			Group:    "moderation_group",
			Consumer: consumerName,
			MinIdle:  30 * time.Second,
			Messages: []string{pend.ID},
		}).Result()

		if err != nil {
			log.Printf("Error claiming message: %v", err)
			continue
		}

		// Process each claimed message
		for _, msg := range claimedMsgs {
			processAndAckMessage(client, msg)
		}
	}
}
