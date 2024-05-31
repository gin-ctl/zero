package route

import (
	"github.com/gin-ctl/zero/app/http/demo/logic/order"
	"github.com/gin-ctl/zero/app/http/demo/logic/user"
	"github.com/gin-gonic/gin"
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

		o := v1.Group("order")
		{
			orderLogic := order.NewOrderLogic()
			p := o.Group("product")
			{
				p.POST("create", orderLogic.Create)
			}

			pay := o.Group("pay")
			{
				pay.POST("create", orderLogic.Create)
			}
		}
	}

}
