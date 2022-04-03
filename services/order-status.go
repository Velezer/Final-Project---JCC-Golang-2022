package services

import (
	"hewantani/models"

	"gorm.io/gorm"
)

type OrderStatus struct {
	Db *gorm.DB
}

func (s OrderStatus) Find(name string) (OrderStatus *models.OrderStatus, err error) {
	err = s.Db.First(&OrderStatus, models.OrderStatus{Name: name}).Error
	if err != nil {
		return nil, err
	}

	return
}
