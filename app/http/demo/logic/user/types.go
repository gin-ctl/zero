package user

import (
	"errors"
	"fmt"
	model "github.com/gin-ctl/zero/model/davari"
	"github.com/gin-ctl/zero/package/database"
	"github.com/gin-ctl/zero/package/http"
	"github.com/gin-ctl/zero/package/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	Id              uint32  `json:"id"`
	Username        string  `json:"username" validate:"required,max=100"`
	Password        string  `json:"password" validate:"required"`
	ConfirmPassword string  `json:"confirm_password" validate:"required"`
	Avatar          string  `json:"avatar" validate:"omitempty,max=255"`
	Phone           *string `json:"phone" validate:"omitempty,phone"`
	Email           string  `json:"email" validate:"omitempty,email"`
	Gender          uint8   `json:"gender" validate:"omitempty,numeric,oneof=0 1 2"`
}

type Index struct {
	Page     uint32  `form:"page" validate:"numeric,min=1"`
	Size     uint32  `form:"size" validate:"numeric,min=1,max=100"`
	Keywords *string `form:"keywords" validate:"omitempty"`
}

func (r *Index) ParseAndCheckParams(c *gin.Context) (err error) {
	err = http.Parse(c, r)
	if err != nil {
		return
	}
	if err = validator.ValidateStructWithOutCtx(r); err != nil {
		return
	}
	return
}

type Create struct {
	User
}

func (r *Create) ParseAndCheckParams(c *gin.Context) (err error) {
	err = http.Parse(c, r)
	if err != nil {
		return
	}
	if err = validator.ValidateStructWithOutCtx(r); err != nil {
		return
	}
	if r.Password != r.ConfirmPassword {
		return errors.New("确认密码与密码不一致")
	}
	err = checkUsername(&r.User, false)
	return
}

type Update struct {
	Id uint32 `path:"id"`
	User
	One *model.User
}

func (r *Update) ParseAndCheckParams(c *gin.Context) (err error) {
	err = http.Parse(c, r)
	if err != nil {
		return
	}
	r.User.Id = r.Id
	err = checkUsername(&r.User, true)
	if err != nil {
		return
	}
	user := model.NewUser()
	err = database.DB.Model(user).Where("id = ?", r.Id).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("用户不存在")
		}
		return
	}
	r.One = user
	return
}

type Info struct {
	Id  uint32 `path:"id"`
	One *model.User
}

func (r *Info) ParseAndCheckParams(c *gin.Context) (err error) {
	err = http.Parse(c, r)
	if err != nil {
		return
	}
	user := model.NewUser()
	err = database.DB.Model(user).Where("id = ?", r.Id).
		Select("id,username,phone,avatar,email,gender,created_at,updated_at").
		First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("用户不存在")
		}
		return
	}
	r.One = user
	return
}

func checkUsername(r *User, isUpdate bool) (err error) {
	user := model.NewUser()
	db := database.DB.Model(user).Where("username = ?", r.Username)
	if isUpdate {
		db = db.Where("id = ?", r.Id)
	}
	// 手机号唯一
	if r.Phone != nil {
		db = db.Where("phone = ?", r.Phone)
	}
	err = db.First(user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if user.Username == r.Username {
		return errors.New(fmt.Sprintf("`%s`已存在", r.Username))
	}
	return
}
