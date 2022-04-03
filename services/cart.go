package services

import (
	"errors"
	"hewantani/models"

	"gorm.io/gorm"
)

type Cart struct {
	Db *gorm.DB
}

func (s Cart) updateCartPrice(cart *models.Cart) {
	cart.TotalPrice = 0
	for _, item := range cart.CartItems {
		cart.TotalPrice += item.Count * item.Product.Price
	}
	s.Db.Save(&cart)
}

func (s Cart) FindAllByuserId(userId uint) (carts *[]models.Cart, err error) {
	err = s.Db.Find(&carts, models.Cart{UserId: userId, IsCheckout: false}).Error
	if err != nil {
		return nil, err
	}
	return
}
func (s Cart) FindById(cartId uint) (cart *models.Cart, err error) {
	err = s.Db.Preload("CartItems.Product").First(&cart, cartId).Error
	if err != nil {
		return nil, err
	}
	s.updateCartPrice(cart)

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

func (s Cart) UpdateCartItem(item *models.CartItem) error {
	if item.Count < 1 {
		return s.Db.Unscoped().Delete(&models.CartItem{}, &item).Error
	}
	return s.Db.Where(&models.CartItem{ProductId: item.ProductId, CartId: item.CartId}).Save(&item).Error
}

func (s Cart) Delete(id uint) (err error) {
	return s.Db.Delete(&models.Cart{}, id).Error
}
