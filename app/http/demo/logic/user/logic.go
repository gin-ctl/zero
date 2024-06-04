package user

import (
	"github.com/gin-ctl/zero/app/http/demo/logic"
	model "github.com/gin-ctl/zero/model/davari"
	"github.com/gin-ctl/zero/package/database"
	"github.com/gin-ctl/zero/package/helper"
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type Logic struct{}

func NewUserLogic() *Logic {
	return &Logic{}
}

func (u *Logic) Index(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Index](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	user := model.NewUser()
	users := make([]model.User, 0)
	var total int64
	db := database.DB.Model(user).
		Select("id,username,phone,avatar,email,gender,created_at,updated_at")
	if params.Data().Keywords != nil {
		db = db.Where("username LIKE ? OR phone LIKE ?", "%"+*params.Data().Keywords+"%", "%"+*params.Data().Keywords+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		http.Alert500(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = db.Find(&users).Error
	if err != nil {
		http.Alert500(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SuccessWithData(c, map[string]interface{}{
		"data":  users,
		"total": total,
	})
}

func (u *Logic) Show(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Info](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	http.SuccessWithData(c, params.Data().One)
}

// Create one row.
func (u *Logic) Create(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Create](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	user := model.NewUser()
	user.Username = params.Data().Username
	user.Password, err = helper.BcryptHash(params.Data().Password)
	if err != nil {
		http.Alert500(c, http.StatusInternalServerError, err.Error())
		return
	}
	user.Phone = cast.ToString(params.Data().Phone)
	user.Avatar = params.Data().Avatar
	user.Gender = params.Data().Gender
	user.Email = params.Data().Email
	err = database.DB.Model(user).Create(user).Error
	if err != nil {
		http.Alert500(c, http.StatusInternalServerError, err.Error())
		return
	}
	http.Success(c)
}

func (u *Logic) Update(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Update](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	params.Data().One.Username = params.Data().Username
	params.Data().One.Phone = *params.Data().Phone
	params.Data().One.Avatar = params.Data().Avatar
	params.Data().One.Email = params.Data().Email
	params.Data().One.Gender = params.Data().Gender
	err = database.DB.Save(params.Data().One).Error
	if err != nil {
		http.Alert500(c, http.StatusInternalServerError, err.Error())
		return
	}
	http.Success(c)
}

func (u *Logic) Destroy(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Info](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	err = database.DB.Delete(params.Data().One).Error
	if err != nil {
		http.Alert500(c, http.StatusInternalServerError, err.Error())
		return
	}
	http.Success(c)
}
