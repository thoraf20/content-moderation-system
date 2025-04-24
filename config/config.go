package config

import (
	"log"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetDefault("REDIS_URL", "localhost:6379")
	viper.SetDefault("STREAM_NAME", "moderation_stream")
	viper.AutomaticEnv()

	log.Printf("Redis: %s | Stream: %s",
		viper.GetString("REDIS_URL"),
		viper.GetString("STREAM_NAME"),
	)
}

func Get(key string) string {
	return viper.GetString(key)
}
