package models

type Cart struct {
	BaseModel

	Name       string `json:"name"`
	IsCheckout bool   `json:"is_checkout" gorm:"default:false"`
	TotalPrice uint   `json:"total_price"`

	UserId uint `json:"user_id"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`

	CartItems []CartItem `json:"cart_items" gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}
