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

func (s Order) FindAllByUserId(userId uint) (m *[]models.Order, err error) {
	err = s.Db.Find(&m, models.Order{UserId: userId}).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s Order) Delete(id uint) (m *models.Order, err error) {
	err = s.Db.Delete(&m, id).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
