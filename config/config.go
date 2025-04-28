package config

import (
	"log"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load from .env first
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found, continuing...")
	}

	viper.AutomaticEnv()
}

func Get(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Printf("Warning: %s not set in config or environment variables", key)
	}
	return value
}
