package models

import (
	"gorm.io/gorm"
)

const (
	ROLE_USER = "USER"
	ROLE_MERCHANT = "MERCHANT"
)

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;unique"`
}
