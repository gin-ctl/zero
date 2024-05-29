package model

import (
    "encoding/json"
	"github.com/gin-ctl/zero/package/time"
)

type Resources struct {
    Id uint32 `json:"id" gorm:"column:id;primaryKey;autoIncrement" validate:"required,numeric"`
    Name string `json:"name" gorm:"column:name" validate:"required,max=255"`
    Type uint8 `json:"type" gorm:"column:type" validate:"required,numeric"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at" validate:"omitempty,datetime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" validate:"omitempty,datetime"`
}


func (p *Resources) TableName() string {
	return "resources"
}

func NewResources() *Resources {
	return &Resources{}
}