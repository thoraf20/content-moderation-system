// cmd/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/thoraf20/api-gateway/internal/config"
	"github.com/thoraf20/api-gateway/internal/router"
	"github.com/thoraf20/api-gateway/internal/middleware"
)

func init() {
	// Load environment variables from config file
	config.LoadConfig()
}

func main() {
	port := viper.GetString("PORT")
	if port == "" {
		port = "8000"
	}

	// Set Gin mode
	if viper.GetString("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())         // Custom logger middleware // Optional rate limiter
	r.Use(middleware.CORSMiddleware()) // Optional CORS

	// Register routes
	router.SetupRoutes(r)

	// Start server
	log.Printf("🚀 API Gateway running on port %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); 
	err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to run server: %v", err)
		os.Exit(1)
	}
}
