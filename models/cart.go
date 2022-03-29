package models

type Cart struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	UserId   uint      `json:"user_id"`
	Products []Product `json:"products" gorm:"many2many:product_cart;"`
}
