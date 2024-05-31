package order

import (
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
)

type Logic struct {
}

func NewOrderLogic() *Logic {
	return &Logic{}
}

func (r *Logic) Create(c *gin.Context) {

	http.Success(c)
}
