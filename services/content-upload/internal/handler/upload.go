package handler

import (
	"content-upload/internal/config"
	"content-upload/internal/event"
	
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form data"})
		return
	}

	files := form.File["files"]
	var uploaded []string

	uploadDir := config.Get("UPLOAD_DIR")
	os.MkdirAll(uploadDir, os.ModePerm)

	for _, file := range files {
		dst := filepath.Join(uploadDir, file.Filename)
		if err := c.SaveUploadedFile(file, dst); 
		err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
			return
		}
		uploaded = append(uploaded, file.Filename)

		// Publish to Kafka (or RabbitMQ later)
		event.PublishModerationEvent(file.Filename, dst)
	}

	c.JSON(http.StatusOK, gin.H{"uploaded": uploaded})
}
