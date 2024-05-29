package model

import (
    "encoding/json"
	"github.com/gin-ctl/zero/package/time"
)

type User struct {
    Id int32 `json:"id" gorm:"column:id;primaryKey;autoIncrement" validate:"required,numeric"`
    Username string `json:"username" gorm:"column:username" validate:"required,max=255"`
    Password string `json:"password" gorm:"column:password" validate:"required,max=255"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at" validate:"omitempty,datetime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" validate:"omitempty,datetime"`
}


func (p *User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}