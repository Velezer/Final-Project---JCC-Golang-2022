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

type StoreIface interface {
	Save(s *models.Store) (*models.Store, error)
}
type ProductIface interface {
	Save(s *models.Product) (*models.Product, error)
}
