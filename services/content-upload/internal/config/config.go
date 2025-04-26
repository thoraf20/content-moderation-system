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

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Read from config.yaml (optional)
	if err := viper.MergeInConfig(); 
	err != nil {
		log.Println("config.yaml not found, continuing...")
	}

	viper.AutomaticEnv() // Allow env vars to override config
}

func Get(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Printf("Warning: %s not set in config or environment variables", key)
	}
	return value
}
