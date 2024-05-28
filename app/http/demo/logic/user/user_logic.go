package user

import (
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
)

type Logic struct {
}

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

func NewUserLogic() *Logic {
	return &Logic{}
}

func (u *Logic) Index(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
		return
	}
	http.SuccessWithData(c, request.Data())
}

func (u *Logic) Show(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
		return
	}
	http.SuccessWithData(c, request.Data())
}

func (u *Logic) Create(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
		return
	}
	http.SuccessWithData(c, request.Data())
}

func (u *Logic) Update(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
		return
	}
	http.SuccessWithData(c, request.Data())
}

func (u *Logic) Destroy(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
		return
	}
	http.SuccessWithData(c, request.Data())
}
