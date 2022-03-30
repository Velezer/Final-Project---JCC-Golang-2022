package services

import (
	"hewantani/models"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Category struct {
	Db *gorm.DB
}

func (s Category) Find(name string) (category *models.Category, err error) {
	err = s.Db.Find(&category, models.Category{Name: name}).Error
	if err != nil {
		return nil, err
	}

	return
}

func (s Category) Save(m *models.Category) (*models.Category, error) {
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))

	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
