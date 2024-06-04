package user

import (
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-ctl/zero/package/validator"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Id        uint64 `path:"id"`
	Page      uint   `query:"page"`
	Size      uint   `query:"size"`
	Name      string `form:"name"`
	FullName  string `json:"name" validate:"required"`
	FolderId  uint64 `json:"folder_id"`
	IsDeleted bool   `json:"is_deleted"`
}

func (r *Request) ParseAndCheckParams(c *gin.Context) (err error) {
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

type Response struct {
	Value string
}

func (r *Response) ParseAndCheckParams(c *gin.Context) (err error) {
	err = http.Parse(c, r)
	if err != nil {
		return
	}
	// TODO: add your logic check.

	return
}
