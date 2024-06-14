package model

import (
	"github.com/gin-ctl/zero/package/time"
)

type AdminUsers struct {
	Id        int32     `json:"id" gorm:"column:id;primaryKey;autoIncrement" validate:"required,numeric"`
	Username  string    `json:"username" gorm:"column:username" validate:"required,max=255"` // 用户名称
	Password  string    `json:"password" gorm:"column:password" validate:"required,max=255"` // 密码
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" validate:"omitempty,datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" validate:"omitempty,datetime"`
}

func (p *AdminUsers) TableName() string {
	return "admin_users"
}

func NewAdminUsers() *AdminUsers {
	return &AdminUsers{}
}
