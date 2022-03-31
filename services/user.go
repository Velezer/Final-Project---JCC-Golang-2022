package services

import (
	"hewantani/models"
	"hewantani/utils/jwttoken"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Db *gorm.DB
}

func (s User) Login(username string, password string) (token string, err error) {
	user := models.User{}
	err = s.Db.Find(&user, models.User{Username: username}).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err = jwttoken.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}
func (s User) Save(u *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Address = html.EscapeString(strings.TrimSpace(u.Address))

	err = s.Db.Create(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s User) ChangePassword(userId uint, password string) (m *models.User, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	m, err = s.FindById(userId)
	if err != nil {
		return nil, err
	}
	m.Password = string(hashedPassword)

	err = s.Db.Model(&m).Updates(&m).Error
	if err != nil {
		return nil, err
	}

	return
}

func (s User) FindById(userId uint) (user *models.User, err error) {
	err = s.Db.Find(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return
}

func (s User) FindByIdJoinRole(userId uint) (user *models.User, err error) {
	err = s.Db.Joins("Role").Find(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return
}
