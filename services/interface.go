package services

import "hewantani/models"

type UserIface interface {
	Save(u *models.User) (*models.User, error)
	Login(username string, password string) (token string, err error)
}

type RoleIface interface {
	Find(name string) (*models.Role, error)
}
