package services

import (
	"hewantani/models"

	"gorm.io/gorm"
)

type Order struct {
	Db *gorm.DB
}

func (s Order) Save(userId, cartId uint) (m *models.Order, err error) {
	m.CartId = cartId
	m.UserId = userId
	err = s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
