package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	UserId uint `json:"user_id"`
	User   User `json:"-"`

	Products []Product `json:"products" gorm:"many2many:product_cart;"`
}
