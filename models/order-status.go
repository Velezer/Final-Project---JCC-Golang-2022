package models

import "gorm.io/gorm"

const (
	ORDER_UNPAID    = "UNPAID"
	ORDER_CANCELLED = "CANCELLED"
	ORDER_COMPLETED = "COMPLETED"
)

type OrderStatus struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;unique"`
}
