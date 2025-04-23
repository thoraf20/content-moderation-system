package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thoraf20/api-gateway/handlers"
	"github.com/thoraf20/api-gateway/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	// Public routes
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// public services
	{
		api.Any("/auth/*path", handlers.ReverseProxy("upload"))
	}

	// Secured services
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth(), middleware.RateLimitMiddleware())
	{
		protected.Any("/upload/*path", handlers.ReverseProxy("upload"))
	}
}
