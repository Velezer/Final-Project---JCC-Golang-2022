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
	m = &models.Order{}
	m.CartId = cartId
	m.UserId = userId

	status, err := All.OrderStatusService.Find(models.ORDER_UNPAID)
	if err != nil {
		return nil, err
	}
	m.Status = *status

	cart, _ := All.CartService.FindById(cartId)
	for _, item := range cart.CartItems {
		user, err := All.UserService.FindById(item.Product.UserId)
		if err != nil {
			return nil, err
		}
		m.Merchants = append(m.Merchants, *user)
	}
	cart.IsCheckout = true
	_, err = All.CartService.Update(cartId, cart)
	if err != nil {
		return nil, err
	}

	err = s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s Order) UpdateStatus(orderId uint, statusName string) (m *models.Order, err error) {
	u := models.Order{}
	status, err := All.OrderStatusService.Find(statusName)
	if err != nil {
		return nil, err
	}
	u.Status = *status

	u.ID = orderId
	err = s.Db.Updates(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s Order) FindAllByUserId(userId uint) (m *[]models.Order, err error) {
	err = s.Db.Joins("Status").Find(&m, models.Order{UserId: userId}).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s Order) FindAllByMerchantId(merchantId uint) (m *[]models.Order, err error) {
	rows, err := s.Db.Table("order_merchant").Select("order_id").Where("user_id", merchantId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	oids := []string{}
	for rows.Next() {
		var orderId string
		if err := rows.Scan(&orderId); err != nil {
			return nil, err
		}

		oids = append(oids, orderId)
	}
	err = s.Db.Joins("Status").Model(&models.Order{}).Where("orders.id", oids).Where("orders.deleted_at", nil).Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s Order) Delete(id uint) (err error) {
	return s.Db.Delete(&models.Order{}, id).Error
}

func (s Order) VerifyOwner(userId uint, found *models.Order) error {
	if found.UserId == userId {
		return nil
	}
	for _, merchant := range found.Merchants {
		if merchant.ID == userId {
			return nil
		}
	}
	return errors.New("this is not your order")
}

func (s Order) FindById(id uint) (Order *models.Order, err error) {
	err = s.Db.Preload("Merchants").Joins("Status").First(&Order, id).Error
	if err != nil {
		return nil, err
	}

	return
}
