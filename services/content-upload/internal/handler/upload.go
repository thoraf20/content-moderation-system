package handler

import (
	"content-upload/internal/config"
	"content-upload/internal/event"
	"fmt"
	"strings"
	"time"

	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func detectFileType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".txt":
		return "text"
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp":
		return "image"
	case ".mp4", ".avi", ".mov", ".mkv":
		return "video"
	default:
		return "unknown"
	}
}

func HandleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form data"})
		return
	}

	var uploaded []string

	uploadDir := config.Get("UPLOAD_DIR")
	os.MkdirAll(uploadDir, os.ModePerm)

	// 1. Handle file uploads if present
	if form != nil && form.File != nil {
		files := form.File["files"]
		for _, file := range files {
			filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
			dst := filepath.Join(uploadDir, filename)

			if err := c.SaveUploadedFile(file, dst); 
			err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
				return
			}

			fileType := detectFileType(file.Filename)
			event.PublishModerationEvent(filename, dst, fileType)
			uploaded = append(uploaded, filename)
		}
	}

	// 2. Handle raw text input
	rawText := c.PostForm("text") // assumes client sends form field named `text`

	if rawText != "" {
		filename := fmt.Sprintf("text_%d", time.Now().UnixNano())
		event.PublishModerationEvent(filename, "", "text", rawText)
		uploaded = append(uploaded, filename)
	}

	c.JSON(http.StatusOK, gin.H{"uploaded": uploaded})
}
