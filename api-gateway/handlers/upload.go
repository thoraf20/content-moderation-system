package handlers

import (
	// "bytes"
	"fmt"
	// "io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func UploadContent(c *gin.Context) {
// 	// Forwarding multipart form to content-upload service
// 	contentUploadServiceURL := "http://content-upload-service:5000/upload"

// 	// Read the multipart form
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form"})
// 		return
// 	}

// 	// Create a proxy request
// 	bodyBuffer := &bytes.Buffer{}
// 	writer := io.MultiWriter(bodyBuffer)

// 	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20) // Max 10MB
// 	c.Request.ParseMultipartForm(10 << 20)

// 	// Send proxy request to content-upload service
// 	req, err := http.NewRequest("POST", contentUploadServiceURL, c.Request.Body)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
// 		return
// 	}
// 	req.Header = c.Request.Header

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "content-upload service unavailable"})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Read and return response from the service
// 	respBody, _ := io.ReadAll(resp.Body)
// 	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
// }

func UploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid form"})
		return
	}

	files := form.File["files"]
	for _, file := range files {
		// Save file locally (just for dev testing)
		err := c.SaveUploadedFile(file, fmt.Sprintf("uploads/%s", file.Filename))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "upload received"})
}
