package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetUrl(r *http.Request, path string) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	return strings.Join([]string{scheme, r.Host, path}, "")
}

// GetRequestIP 获取ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
