package services

import (
	"hewantani/models"
)

type UserIface interface {
	Save(u *models.User) (*models.User, error)
	Login(username string, password string) (token string, err error)
	FindById(userId uint) (*models.User, error)
}

type RoleIface interface {
	Find(name string) (*models.Role, error)
}
type CategoryIface interface {
	Find(name string) (*models.Category, error)
	Save(s *models.Category) (*models.Category, error)
}

type StoreIface interface {
	Save(s *models.Store) (*models.Store, error)
}
type CartIface interface {
	Save(s *models.Cart) (*models.Cart, error)
	AddCartItem(cartId uint, item models.CartItem) (m *models.Cart, err error)
}
type ProductIface interface {
	FindById(id uint) (*models.Product, error)
	Save(s *models.Product) (*models.Product, error)
}
