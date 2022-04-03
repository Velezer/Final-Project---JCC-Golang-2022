package models

type CartItem struct {
	BaseModel

	ProductId uint    `json:"product_id" gorm:"uniqueIndex:product_cart_pair;"`
	Product   Product `json:"product,omitempty"`

	Count uint `json:"count"`

	CartId uint `json:"-" gorm:"uniqueIndex:product_cart_pair;"`
}
