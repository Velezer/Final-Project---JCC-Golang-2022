package controllers

import (
	"fmt"
	"net/http"

	"hewantani/models"
	"hewantani/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=USER MERCHANT"`
}

// Register godoc
// @Summary      Register a user.
// @Description  registering a user from public access.
// @Tags         Auth
// @Param        Body  body  RegisterInput  true  "the body to register a user"
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var roleService services.RoleIface = services.Role{Db: db}
	role, err := roleService.Find(input.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	u.Address = input.Address
	u.Role = *role

	var userService services.UserIface = services.User{Db: db}
	savedUser, err := userService.Save(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": savedUser.Username,
		"email":    savedUser.Email,
		"address":  savedUser.Address,
		"role":     savedUser.Role.Name,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})

}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access api by user's role.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userService services.UserIface = services.User{Db: db}
	token, err := userService.Login(input.Username, input.Password)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	data := map[string]string{
		"token": token,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})

}
