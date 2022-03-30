package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model

	Name        string `json:"name" gorm:"not null;unique"`
	Description string `json:"description"`
	Address     string `json:"address"`

	UserId uint `json:"user_id"`
	User   User `json:"-"`

	Products []Product `json:"-"`
}