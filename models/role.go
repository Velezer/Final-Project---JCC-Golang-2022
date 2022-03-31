package models

import (
	"gorm.io/gorm"
)

const (
	ROLE_USER     = "USER"
	ROLE_MERCHANT = "MERCHANT"
	ROLE_ADMIN    = "ADMIN"
)

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;unique"`
}
