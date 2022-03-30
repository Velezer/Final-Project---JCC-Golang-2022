package services

import (
	"hewantani/models"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	Db *gorm.DB
}

func (s Product) Save(m *models.Product) (*models.Product, error) {
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))

	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	// err = s.Db.Save(&m).Error
	// if err != nil {
	// 	return nil, err
	// }

	return m, nil
}
