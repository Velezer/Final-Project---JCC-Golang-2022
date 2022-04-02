package services

import (
	"hewantani/models"

	"gorm.io/gorm"
)

type Role struct {
	Db *gorm.DB
}

func (s Role) Find(name string) (role *models.Role, err error) {
	err = s.Db.First(&role, models.Role{Name: name}).Error
	if err != nil {
		return nil, err
	}

	return
}
