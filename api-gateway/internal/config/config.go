package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load from .env first
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // Try to load .env
	if err != nil {
		log.Println("No .env file found, continuing...")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	// Read from config.yaml (optional)
	if err := viper.MergeInConfig(); err != nil {
		log.Println("config.yaml not found, continuing...")
	}

	viper.AutomaticEnv() // Allow env vars to override config

	LoadServiceMap()
}
