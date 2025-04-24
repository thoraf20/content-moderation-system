package main

import (
	"log"
	"content-analysis/config"
	// "content-analysis/consumer"
)

func main() {
	config.LoadConfig() // loads from env or config file

	log.Println("Starting Content Moderation Service...")

	// Start consuming from Redis Stream
	// consumer.StartStreamConsumer()
}
