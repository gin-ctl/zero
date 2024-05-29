package model

import (
	"github.com/gin-ctl/zero/package/time"
)

type Policies struct {
	Id         uint32    `json:"id" gorm:"column:id;primaryKey;autoIncrement" validate:"omitempty,numeric"`
	RoleId     uint32    `json:"role_id" gorm:"column:role_id" validate:"omitempty,numeric"`
	ResourceId uint32    `json:"resource_id" gorm:"column:resource_id" validate:"omitempty,numeric"`
	Action     string    `json:"action" gorm:"column:action" validate:"omitempty,max=6"`
	Effect     string    `json:"effect" gorm:"column:effect" validate:"omitempty,max=5"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" validate:"omitempty,datetime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" validate:"omitempty,datetime"`
}

func (p *Policies) TableName() string {
	return "policies"
}

func NewPoliciesStruct() *Policies {
	return &Policies{}
}
