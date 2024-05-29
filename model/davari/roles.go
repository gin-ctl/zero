package model

import (
  	"github.com/gin-ctl/zero/package/time"
)

type Roles struct {
    Id uint32 `json:"id" gorm:"column:id;primaryKey;autoIncrement" validate:"required,numeric"`
    Name string `json:"name" gorm:"column:name" validate:"required,max=255"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at" validate:"omitempty,datetime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" validate:"omitempty,datetime"`
}


func (p *Roles) TableName() string {
	return "roles"
}

func NewRoles() *Roles {
	return &Roles{}
}