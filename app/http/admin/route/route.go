package route

import (
	"github.com/gin-ctl/zero/app/http/admin/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAdminAPI(r *gin.Engine) {

	// middlewares
	r.Use(
		middleware.Auth(),
	)

}
