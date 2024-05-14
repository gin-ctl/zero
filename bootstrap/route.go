package bootstrap

import (
	"github.com/gin-gonic/gin"
	"zero/app/http/demo/route"
	"zero/app/middleware"
	"zero/package/http"
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
