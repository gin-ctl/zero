package model



type UserRoles struct {
    UserId uint32 `json:"user_id" gorm:"column:user_id"`
    RoleId uint32 `json:"role_id" gorm:"column:role_id"`
}


func (p *UserRoles) TableName() string {
	return "user_roles"
}

func NewUserRoles() *UserRoles {
	return &UserRoles{}
}