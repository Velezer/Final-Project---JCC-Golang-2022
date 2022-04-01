package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"not null;unique"`
	Username string `json:"username" gorm:"unique;not null;type:varchar(255);check:username <> ''"`
	Password string `json:"password"`
	Address  string `json:"address"`

	RoleId uint `json:"role_id" gorm:"not null"`
	Role   Role `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
