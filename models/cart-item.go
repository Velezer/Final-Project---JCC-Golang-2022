package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model

	ProductId uint    `json:"product_id"`
	Product   Product `json:"product"`

	Count uint `json:"count"`

	CartId uint `json:"cart_id"`
	Cart   Cart `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
