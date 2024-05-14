package base

import "zero/package/time"

type PrimaryKey struct {
	Id uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,optional"`
}

type CommonTime struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,optional"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,optional"`
}

type CreatedAt struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,optional"`
}
