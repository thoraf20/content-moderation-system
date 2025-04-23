package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/thoraf20/api-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func ReverseProxy(service string) gin.HandlerFunc {
	target := config.ServiceMap[service]
	return func(c *gin.Context) {
		destination, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(destination)
		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/"+service)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
