package services

import (
	"hewantani/models"
)

type UserIface interface {
	Save(u *models.User) (*models.User, error)
	Login(username string, password string) (token string, err error)
	FindById(userId uint) (*models.User, error)
	FindByIdJoinRole(userId uint) (*models.User, error)
	Delete(id uint) error
	ChangePassword(id uint, password string) error
	Update(userId uint, u *models.User) (m *models.User, err error)
}

type RoleIface interface {
	Find(name string) (*models.Role, error)
}
type OrderStatusIface interface {
	Find(name string) (*models.OrderStatus, error)
}
type CategoryIface interface {
	Find(name string) (*models.Category, error)
	Save(s *models.Category) (*models.Category, error)
}

type StoreIface interface {
	Save(s *models.Store) (*models.Store, error)
	FindAll(keyword string) (*[]models.Store, error)
	FindById(id uint) (*models.Store, error)
	Update(id uint, u *models.Store) (*models.Store, error)
	Delete(id uint) error
	VerifyOwner(userId uint, found *models.Store) error
}
type CartIface interface {
	Save(s *models.Cart) (*models.Cart, error)
	Update(id uint, u *models.Cart) (*models.Cart, error)
	Delete(id uint) error

	UpdateCartItem(item *models.CartItem) (err error)

	FindAllByuserId(userId uint, isCheckout bool) (m *[]models.Cart, err error)
	FindById(id uint) (m *models.Cart, err error)
	VerifyOwner(userId uint, found *models.Cart) error
}
type ProductIface interface {
	FindById(id uint) (*models.Product, error)
	FindAll(categories []string, keyword string) (*[]models.Product, error)
	Save(s *models.Product) (*models.Product, error)
	Update(id uint, u *models.Product) (*models.Product, error)
	Delete(id uint) error
	VerifyOwner(userId uint, found *models.Product) error
}
type OrderIface interface {
	FindById(id uint) (*models.Order, error)
	VerifyOwner(userId uint, found *models.Order) error
	Save(userId, cartId uint) (*models.Order, error)
	UpdateStatus(orderId uint, statusName string) (*models.Order, error)
	FindAllByUserId(userId uint) (*[]models.Order, error)
	Delete(id uint) error
	FindAllByMerchantId(merchantId uint) (m *[]models.Order, err error)
}
