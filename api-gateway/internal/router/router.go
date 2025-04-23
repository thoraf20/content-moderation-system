package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thoraf20/api-gateway/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/upload", handlers.UploadHandler)
	}
}
