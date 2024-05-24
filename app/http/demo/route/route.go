package route

import (
	"github.com/gin-gonic/gin"
	"zero/app/http/demo/logic/user"
)

func RegisterDemoAPI(r *gin.Engine) {

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
