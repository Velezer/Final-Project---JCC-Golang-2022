package models

const (
	ORDER_UNPAID    = "UNPAID"
	ORDER_CANCELLED = "CANCELLED"
	ORDER_PAID      = "PAID"
	ORDER_SHIPPING  = "SHIPPING"
	ORDER_DELIVERED = "DELIVERED"
)

var orderState map[string]int = map[string]int{
	ORDER_CANCELLED: -1,
	ORDER_UNPAID:    0,
	ORDER_PAID:      1,
	ORDER_SHIPPING:  2,
	ORDER_DELIVERED: 3,
}

func (m OrderStatus) GetState(status string) int {
	return orderState[status]
}

type OrderStatus struct {
	BaseModel
	Name string `json:"name" gorm:"not null;unique"`
}

const (
	ROLE_USER     = "USER"
	ROLE_MERCHANT = "MERCHANT"
)

type Role struct {
	BaseModel
	Name string `json:"name" gorm:"not null;unique"`
}
