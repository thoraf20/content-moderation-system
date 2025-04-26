package config

import "github.com/spf13/viper"

var ServiceMap map[string]string

func LoadServiceMap() {
	ServiceMap = map[string]string{
		"auth":   viper.GetString("SERVICE_AUTH_URL"),
		"upload": viper.GetString("SERVICE_UPLOAD_URL"),
		// Add more services here
	}
}
