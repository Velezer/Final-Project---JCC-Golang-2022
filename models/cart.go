package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	Name string `json:"name"`

	UserId uint `json:"user_id"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TotalPrice uint `json:"total_price"`

	CartItems []CartItem `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
