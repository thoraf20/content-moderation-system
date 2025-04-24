package handler

import (
	"content-upload/internal/config"
	"content-upload/internal/event"
	"fmt"
	"time"

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
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		dst := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, dst); 
		err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
			return
		}

		uploaded = append(uploaded, filename)
		event.PublishModerationEvent(filename, dst)
	}

	c.JSON(http.StatusOK, gin.H{"uploaded": uploaded})
}
