package model

import "time"

type UserRole struct {  //==> user_roles
	ID        int `gorm:"primaryKey"`
	RoleID    int `gorm:"index"`
	UserID    int `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (UserRole) TableName() string {
	return "user_role"
}