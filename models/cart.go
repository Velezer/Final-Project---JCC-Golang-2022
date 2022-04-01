package models

type Cart struct {
	BaseModel

	Name string `json:"name"`

	UserId uint `json:"user_id"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`

	TotalPrice uint `json:"total_price"`

	CartItems []CartItem `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
