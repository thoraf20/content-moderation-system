package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thoraf20/api-gateway/handlers"
	"github.com/thoraf20/api-gateway/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// public services
	{
		api.Any("/auth/*path", handlers.ReverseProxy("upload"))
	}

	// Secured services
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		protected.Any("/upload/*path", handlers.ReverseProxy("upload"))
	}
}
