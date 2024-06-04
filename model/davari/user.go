package model

import (
  	"github.com/gin-ctl/zero/package/time"
)

type User struct {
    Id int32 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
    Username string `json:"username" gorm:"column:username"`
    Password string `json:"password" gorm:"column:password"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}


func (p *User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}