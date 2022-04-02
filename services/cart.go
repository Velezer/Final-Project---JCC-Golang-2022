package services

import (
	"errors"
	"hewantani/models"

	"gorm.io/gorm"
)

type Cart struct {
	Db *gorm.DB
}

func (s Cart) FindByuserId(userId uint) (cart *models.Cart, err error) {
	err = s.Db.First(&cart, models.Cart{UserId: userId, IsCheckout: false}).Error
	if err != nil {
		return nil, err
	}

	err = s.Db.Model(cart).Association("CartItems").Find(&cart.CartItems)
	if err != nil {
		return nil, err
	}

	return
}
func (s Cart) FindById(cartId uint) (cart *models.Cart, err error) {
	err = s.Db.First(&cart, cartId).Error
	if err != nil {
		return nil, err
	}

	return
}
func (s Cart) VerifyOwner(userId uint, found *models.Cart) error {
	if found.UserId == userId {
		return nil
	}
	return errors.New("this is not your cart")
}
func (s Cart) Save(m *models.Cart) (*models.Cart, error) {
	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s Cart) Update(cartId uint, m *models.Cart) (*models.Cart, error) {
	m.ID = cartId
	err := s.Db.Updates(&m).Error
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
	err = s.Db.Model(&m).Save(&m).Error
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
