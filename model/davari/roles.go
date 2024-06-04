package model

import (
  	"github.com/gin-ctl/zero/package/time"
)

type Roles struct {
    Id uint32 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
    Name string `json:"name" gorm:"column:name"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}


func (p *Roles) TableName() string {
	return "roles"
}

func NewRoles() *Roles {
	return &Roles{}
}