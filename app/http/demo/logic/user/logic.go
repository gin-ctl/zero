package user

import (
	"github.com/gin-ctl/zero/app/http/demo/logic"
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
)

type Logic struct{}

func NewUserLogic() *Logic {
	return &Logic{}
}

func (u *Logic) Index(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Request](c)
	if err != nil {
		http.Fail(c, err)
		return
	}
	//id := params.Data().Id
	http.SuccessWithData(c, params.Data())
}

func (u *Logic) Show(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Response](c)
	if err != nil {
		http.Fail(c, err)
		return
	}

	http.SuccessWithData(c, params.Data())
}

func (u *Logic) Create(c *gin.Context) {
	http.Success(c)
}

func (u *Logic) Update(c *gin.Context) {
	http.Success(c)
}

func (u *Logic) Destroy(c *gin.Context) {
	http.Success(c)
}
