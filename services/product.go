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

	return m, nil
}

func (s Product) Update(id uint, u *models.Product) (m *models.Product, err error) {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	err = s.Db.Find(&m, id).Updates(&u).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s Product) Delete(id uint) (m *models.Product, err error) {
	err = s.Db.Delete(&m, id).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s Product) FindById(id uint) (product *models.Product, err error) {
	err = s.Db.Find(&product, id).Error
	if err != nil {
		return nil, err
	}

	return
}

func (s Product) FindAll() (products *[]models.Product, err error) {
	err = s.Db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return
}
