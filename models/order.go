package models

type Order struct {
	BaseModel

	StatusId uint        `json:"status_id"`
	Status   OrderStatus `json:"status"`

	UserId uint `json:"user_id"`
	User   User `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`

	Merchants []User `json:"-" gorm:"many2many:order_merchant;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`

	CartId uint `json:"cart_id"`
	Cart   Cart `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
