package user

import (
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-ctl/zero/package/validator"
	"github.com/gin-gonic/gin"
)

// Demo demo.
//type Demo struct {
//	Page     uint32  `form:"page" validate:"numeric,min=1"`
//	Size     uint32  `form:"size" validate:"numeric,min=1,max=100"`
//	Keywords *string `form:"keywords" validate:"omitempty"`
//}
//
//func (r *Demo) ParseAndCheckParams(c *gin.Context) (err error) {
//	err = http.Parse(c, r)
//	if err != nil {
//		return
//	}
//	if err = validator.ValidateStructWithOutCtx(r); err != nil {
//		return
//	}
//	return
//}

type Ping struct {
    // TODO: add your params.
}

func (r *Ping) ParseAndCheckParams(c *gin.Context) (err error) {
	err = http.Parse(c, r)
	if err != nil {
		return
	}
	if err = validator.ValidateStructWithOutCtx(r); err != nil {
		return
	}
	// TODO: add your logic check.
	return
}
