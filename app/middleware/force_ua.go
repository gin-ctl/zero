package middlewares

import (
	"github.com/gin-gonic/gin"
	"zero/package/http"
)

// ForceUA 中间件，强制请求头部必须附带 User-Agent
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.UserAgent() == "" {
			http.Alert400WithoutMessage(c, http.MissUserAgent)
		}
		c.Next()
	}
}
