package services

import (
	"hewantani/models"

	"gorm.io/gorm"
)

type Cart struct {
	Db *gorm.DB
}

func (s Cart) FindByuserId(userId uint) (cart *models.Cart, err error) {
	err = s.Db.Find(&cart, models.Cart{UserId: userId}).Error
	if err != nil {
		return nil, err
	}

	return
}
func (s Cart) FindById(cartId uint) (cart *models.Cart, err error) {
	err = s.Db.Find(&cart, cartId).Error
	if err != nil {
		return nil, err
	}

	return
}

func (s Cart) Save(m *models.Cart) (*models.Cart, error) {
	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s Cart) AddCartItem(cartId uint, item models.CartItem) (m *models.Cart, err error) {
	m, err = s.FindById(cartId)
	if err != nil {
		return nil, err
	}

	m.CartItems = append(m.CartItems, item)
	err = s.Db.Model(&m).Updates(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s Cart) DeleteCartItem(itemId uint) (m *models.Cart, err error) {
	err = s.Db.Model(&m).Delete(itemId).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
