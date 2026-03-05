package user

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleSuperUser Role = "superuser"
)

type User struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Username  string        `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string        `gorm:"size:255;not null" json:"-"`
	Role      Role          `gorm:"size:20;default:superuser" json:"role"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) TableName() string {
	return "user"
}
