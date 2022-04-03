package services

import (
	"errors"
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

func (s Store) FindAll(keyword string) (m *[]models.Store, err error) {
	if len(keyword) > 0 {
		err = s.Db.Where("name like ?", "%"+keyword+"%").Find(&m).Error
	} else {
		err = s.Db.Find(&m).Error
	}
	if err != nil {
		return nil, err
	}

	return
}
func (s Store) FindById(id uint) (m *models.Store, err error) {
	err = s.Db.First(&m, id).Error
	if err != nil {
		return nil, err
	}

	return
}
func (s Store) VerifyOwner(userId uint, found *models.Store) error {
	if found.UserId == userId {
		return nil
	}
	return errors.New("this is not your store")
}

func (s Store) Update(id uint, u *models.Store) (m *models.Store, err error) {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	err = s.Db.Find(&m, id).Updates(&u).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s Store) Delete(id uint) (err error) {
	err = s.Db.Delete(&models.Store{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
