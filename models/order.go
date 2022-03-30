package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	UserId uint `json:"user_id"`
	User   User `json:"-"`

	CartId uint `json:"cart_id"`
	Cart   Cart `json:"-"`
}
