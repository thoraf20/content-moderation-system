package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		log := map[string]interface{}{
			"status":   status,
			"method":   method,
			"path":     path,
			"duration": duration.String(),
		}
		// You can plug into a real logger here
		fmt.Println(log)
	}
}
