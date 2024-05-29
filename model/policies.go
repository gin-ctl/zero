package model

import (
    "encoding/json"
	"github.com/gin-ctl/zero/package/time"
)

type Policies struct {
    Id uint32 `json:"id" gorm:"column:id;primaryKey;autoIncrement" validate:"required,numeric"`
    RoleId uint32 `json:"role_id" gorm:"column:role_id" validate:"required,numeric"`
    ResourceId uint32 `json:"resource_id" gorm:"column:resource_id" validate:"required,numeric"`
    Action string `json:"action" gorm:"column:action" validate:"required,max=6"`
    Effect string `json:"effect" gorm:"column:effect" validate:"required,max=5"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at" validate:"omitempty,datetime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" validate:"omitempty,datetime"`
}


func (p *Policies) TableName() string {
	return "policies"
}

func NewPolicies() *Policies {
	return &Policies{}
}