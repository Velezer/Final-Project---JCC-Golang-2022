package services

import (
	"errors"
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

func (s Product) Delete(id uint) (err error) {
	return s.Db.Delete(&models.Product{}, id).Error
}

func (s Product) FindById(id uint) (product *models.Product, err error) {
	err = s.Db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return
}
func (s Product) VerifyOwner(userId uint, found *models.Product) error {
	if found.UserId == userId {
		return nil
	}
	return errors.New("this is not your product")
}
func (s Product) FindAll(categories []string) (products *[]models.Product, err error) {
	if len(categories) > 0 {
		cats := models.Category{}
		err = s.Db.Preload("Products").Model(&models.Category{}).Where("categories.name", categories).Find(&cats).Error
		products = &cats.Products
	} else {
		err = s.Db.Find(&products).Error
	}
	if err != nil {
		return nil, err
	}

	return
}
