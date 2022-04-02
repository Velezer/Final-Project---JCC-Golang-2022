package models

type Order struct {
	BaseModel

	StatusId uint        `json:"status_id"`
	Status   OrderStatus `json:"-"`

	UserId uint `json:"user_id"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`

	MerchantId uint `json:"merchant_id"`
	Merchant   User `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`

	CartId uint `json:"cart_id"`
	Cart   Cart `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
