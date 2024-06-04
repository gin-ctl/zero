package model

import (
  	"github.com/gin-ctl/zero/package/time"
)

type AdminUsers struct {
    Id int32 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
    Username string `json:"username" gorm:"column:username"`
    Password string `json:"password" gorm:"column:password"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}


func (p *AdminUsers) TableName() string {
	return "admin_users"
}

func NewAdminUsers() *AdminUsers {
	return &AdminUsers{}
}