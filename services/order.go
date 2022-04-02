package services

import (
	"errors"
	"hewantani/models"

	"gorm.io/gorm"
)

type Order struct {
	Db *gorm.DB
}

func (s Order) Save(userId, cartId uint) (m *models.Order, err error) {
	m.CartId = cartId
	m.UserId = userId
	m.Status.Name = models.ORDER_UNPAID
	err = s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s Order) UpdateStatus(orderId uint, statusName string) (m *models.Order, err error) {
	u := models.Order{}
	u.Status.Name = statusName
	err = s.Db.Model(&models.Order{}).Updates(&u).Error
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

func (s Order) Delete(id uint) (err error) {
	return s.Db.Delete(&models.Order{}, id).Error
}

func (s Order) VerifyOwner(userId uint, found *models.Order) error {
	if found.UserId == userId || found.MerchantId == userId {
		return nil
	}
	return errors.New("this is not your order")
}

func (s Order) FindById(id uint) (Order *models.Order, err error) {
	err = s.Db.First(&Order, id).Error
	if err != nil {
		return nil, err
	}

	return
}
