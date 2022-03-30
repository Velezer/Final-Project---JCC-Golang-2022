package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"not null;unique"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password"`
	Address  string `json:"address"`

	RoleId uint `json:"role_id" gorm:"not null"`
	Role   Role `json:"-"`
}
