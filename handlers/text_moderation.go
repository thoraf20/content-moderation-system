package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"content-analysis/moderation"
)

func HandleTextModeration(c *gin.Context) {
	var req struct {
		Text string `json:"text"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := moderation.ModerateText(req.Text)
	if result.Flagged {
		c.JSON(http.StatusOK, gin.H{"moderated": false, "reason": result.Reason, "offenses": result.Offenses})
		return
	}

	c.JSON(http.StatusOK, gin.H{"moderated": true, "message": "Content is clean"})
}
