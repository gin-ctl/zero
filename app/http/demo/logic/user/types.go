package user

import (
	"github.com/gin-ctl/zero/package/http"
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
	Dive      Dive   `json:"dive"`
}

func (r *Request) ParseAndCheckParams(c *gin.Context) (err error) {
	err = http.Parse(c, r)
	if err != nil {
		return
	}
	// todo 业务校验

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
	// todo 业务校验

	return
}

type Dive struct {
	UserInfo Info `json:"user_info"`
}

type Info struct {
	Name  string `json:"name"`
	Phone uint   `json:"phone"`
	Fans  []Fans `json:"fans"`
}

type Fans struct {
	Name  string `json:"name"`
	Phone uint   `json:"phone"`
}
