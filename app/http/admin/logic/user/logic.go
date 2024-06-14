package user

import (
	"github.com/gin-ctl/zero/app/http/admin/logic"
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
)

type Logic struct{}

func NewLogic() *Logic {
	return &Logic{}
}

// Ping ping
func (l *Logic) Ping(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Ping](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: Your logic.

	// TODO: Replace your return struct.
	http.SuccessWithData(c, params.Data())
}
