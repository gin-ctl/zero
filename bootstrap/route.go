package bootstrap

import (
	"github.com/gin-ctl/zero/app/http/demo/route"

	// {{.Import}}
	"github.com/gin-ctl/zero/app/middleware"
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.Cors(),
		middlewares.ForceUA(),
		// register global middleware.
		// {{.GlobalMiddleware}}
	)
}

func RegisterDemoApiRoute(router *gin.Engine) {
	// route not found.
	http.Alert404Route(router)
	// global middleware.
	RegisterGlobalMiddleware(router)
	// Initialize route.
	route.RegisterDemoAPI(router)
}

// {{.ApiRoute}}
