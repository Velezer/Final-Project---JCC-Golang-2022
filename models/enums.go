package models

const (
	ORDER_UNPAID    = "UNPAID"
	ORDER_CANCELLED = "CANCELLED"
	ORDER_COMPLETED = "COMPLETED"
)

type OrderStatus struct {
	BaseModel
	Name string `json:"name" gorm:"not null;unique"`
}

const (
	ROLE_USER     = "USER"
	ROLE_MERCHANT = "MERCHANT"
	ROLE_ADMIN    = "ADMIN"
)

type Role struct {
	BaseModel
	Name string `json:"name" gorm:"not null;unique"`
}
