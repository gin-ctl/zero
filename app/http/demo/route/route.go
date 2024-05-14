package route

import (
	"github.com/gin-gonic/gin"
	"zero/package/http"
)

func RegisterDemoAPI(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/user", func(c *gin.Context) {
			http.Success(c)
		})
	}
}
