package services

import (
	"hewantani/models"
	"hewantani/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_User_Register(t *testing.T) {
	db := tests.SetupDB()
	Setup(db)
	assert := assert.New(t)

	t.Run("valid user register", func(t *testing.T) {
		u := models.User{}
		u.Email = "user@email.com"
		u.Username = "username"
		u.Password = "password"
		u.Address = "address"

		saved, err := All.UserService.Register(&u)
		assert.NoError(err)

		assert.Equal(u.Email, saved.Email)
		assert.Equal(u.Username, saved.Username)
		assert.Equal(u.Password, saved.Password)
		assert.Equal(u.Address, saved.Address)
		assert.Equal(&u, &u)
	})

}
