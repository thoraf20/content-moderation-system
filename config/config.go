package config

import (
	"log"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetDefault("BROKER_URL", "localhost:6379")
	viper.SetDefault("TOPIC_NAME", "content_events")
	viper.AutomaticEnv()

	log.Printf("Redis: %s | Stream: %s",
		viper.GetString("BROKER_URL"),
		viper.GetString("TOPIC_NAME"),
	)
}

func Get(key string) string {
	return viper.GetString(key)
}
