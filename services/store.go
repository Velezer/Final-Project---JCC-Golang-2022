package services

import (
	"hewantani/models"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Store struct {
	Db *gorm.DB
}

func (s Store) Save(m *models.Store) (*models.Store, error) {
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))
	m.Description = html.EscapeString(strings.TrimSpace(m.Description))
	m.Address = html.EscapeString(strings.TrimSpace(m.Address))

	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
