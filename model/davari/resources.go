package model

import (
  	"github.com/gin-ctl/zero/package/time"
)

type Resources struct {
    Id uint32 `json:"id" gorm:"column:id;primaryKey;autoIncrement"` 
    Name string `json:"name" gorm:"column:name"`	// 资源名称 
    Type uint8 `json:"type" gorm:"column:type"`	// 资源类型 
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` 
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` 
}


func (p *Resources) TableName() string {
	return "resources"
}

func NewResources() *Resources {
	return &Resources{}
}