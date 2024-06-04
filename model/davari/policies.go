package model

import (
  	"github.com/gin-ctl/zero/package/time"
)

type Policies struct {
    Id uint32 `json:"id" gorm:"column:id;primaryKey;autoIncrement"` 
    RoleId uint32 `json:"role_id" gorm:"column:role_id"`	// 角色id 
    ResourceId uint32 `json:"resource_id" gorm:"column:resource_id"`	// 资源id 
    Action string `json:"action" gorm:"column:action"`	// 操作 
    Effect string `json:"effect" gorm:"column:effect"`	// 影响：ALLOW-允许、DENY-拒绝 
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` 
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` 
}


func (p *Policies) TableName() string {
	return "policies"
}

func NewPolicies() *Policies {
	return &Policies{}
}