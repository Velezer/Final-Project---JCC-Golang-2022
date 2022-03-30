package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	Name string `json:"name"`

	UserId uint `json:"user_id"`
	User   User `json:"-"`

	TotalPrice uint `json:"total_price"`

	CartItems []CartItem `json:"-"`
}
