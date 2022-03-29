package models

type Order struct {
	Id     uint `json:"id" gorm:"primary_key"`
	UserId uint `json:"user_id"`

	CartId uint `json:"cart_id"`
	Cart   Cart `json:"-"`
}
