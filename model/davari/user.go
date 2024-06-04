package model

import (
	"github.com/gin-ctl/zero/package/time"
)

type User struct {
	Id        int32     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"column:username"` // 用户名称
	Password  string    `json:"password" gorm:"column:password"` // 密码
	Avatar    string    `json:"avatar" gorm:"column:avatar"`     // 头像
	Phone     string    `json:"phone" gorm:"column:phone"`       // 电话号码
	Email     string    `json:"email" gorm:"column:email"`       // 邮箱
	Gender    uint8     `json:"gender" gorm:"column:gender"`     // 性别：0-未知、1-男、2-女
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (p *User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}
