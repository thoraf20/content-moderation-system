package main

import (
	"content-upload/internal/config"
	"content-upload/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig() // Viper config

	r := gin.Default()
	r.POST("/upload", handler.HandleUpload)

	port := config.Get("PORT")
	if port == "" {
		port = "5001"
	}
	r.Run(":" + port)
}
