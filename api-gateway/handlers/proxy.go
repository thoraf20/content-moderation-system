package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ReverseProxy(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		target := viper.GetString("services." + serviceName)
		if target == "" {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Unknown service"})
			return
		}

		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/api/"+serviceName)
		c.Request.Host = remote.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
