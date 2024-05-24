package user

import (
	"github.com/gin-gonic/gin"
	"zero/package/http"
)

type Logic struct {
}

type Request struct {
	Id   uint64 `path:"id"`
	Page uint   `query:"page"`
	Size uint   `query:"size"`
	Name string `form:"name"`
	//FullName  string `json:"name"`
	FolderId  uint64 `json:"folder_id"`
	IsDeleted bool   `json:"is_deleted"`
}

func NewUserLogic() *Logic {
	return &Logic{}
}

func (u *Logic) Index(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
	}
	http.SuccessWithData(c, request.Data())
}

func (u *Logic) Show(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
	}
	http.SuccessWithData(c, request.Data())
}

func (u *Logic) Create(c *gin.Context) {
	var req Request
	request, err := http.Parse(c, req)
	if err != nil {
		http.Fail(c, err)
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
	}
	http.SuccessWithData(c, request.Data())
}
