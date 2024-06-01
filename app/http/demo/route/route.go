package route

import (
	"github.com/gin-ctl/zero/app/http/demo/logic/user"
	"github.com/gin-ctl/zero/app/http/demo/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterDemoAPI(r *gin.Engine) {

	// middlewares
	r.Use(
		// {{.Middleware}}
		middleware.Auth(),
	)

	v1 := r.Group("v1")
	{
		// 用户
		u := v1.Group("user")
		{
			userLogic := user.NewUserLogic()
			u.GET("", userLogic.Index)
			u.GET(":id", userLogic.Show)
			u.POST("", userLogic.Create)
			u.PUT(":id", userLogic.Update)
			u.DELETE(":id", userLogic.Destroy)
		}
	}

}
