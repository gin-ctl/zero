package bootstrap

import (
	admin "github.com/gin-ctl/zero/app/http/admin/route"
	demo "github.com/gin-ctl/zero/app/http/demo/route"
	"github.com/gin-ctl/zero/middleware"
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddleware(r *gin.Engine) {
	r.Use(
		middlewares.Auth(),
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.Cors(),
		middlewares.ForceUA(),
	)
}

func RegisterDemoApiRoute(router *gin.Engine) {
	// route not found.
	http.Alert404Route(router)
	// global middleware.
	RegisterGlobalMiddleware(router)
	// Initialize route.
	demo.RegisterDemoAPI(router)
}

func RegisterAdminApiRoute(router *gin.Engine) {
	// route not found.
	http.Alert404Route(router)
	// global middleware.
	RegisterGlobalMiddleware(router)
	// Initialize route.
	admin.RegisterAdminAPI(router)
}
