package controllers

import (
	"net/http"
	"strconv"

	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

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
// @Tags         User
// @Param        Body  body  RegisterInput  true  "the body to register a user"
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /users [post]
func (a UserController) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	role, err := services.All.RoleService.Find(input.Role)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	u.Address = input.Address
	u.Role = *role

	savedUser, err := services.All.UserService.Save(&u)
	if err != nil {
		c.Error(err)
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

type changePasswordInput struct {
	Password string `json:"password" binding:"required"`
}

// Change Password godoc
// @Summary      change user's password
// @Description  change user's password
// @Tags         User
// @Param        Body  body  changePasswordInput  true  "the body to register a user"
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /users/:id/password [post]
func (a UserController) ChangePassword(c *gin.Context) {
	var input changePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	userIdString := c.Param("id")
	userId, err := strconv.ParseUint(userIdString, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	savedUser, err := services.All.UserService.ChangePassword(uint(userId), input.Password)
	if err != nil {
		c.Error(err)
		return
	}

	user := map[string]string{
		"username": savedUser.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed success", "data": user})

}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access api by user's role.
// @Tags User
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/login [post]
func (a UserController) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	token, err := services.All.UserService.Login(input.Username, input.Password)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest).SetData("username or password is incorrect"))
		return
	}

	data := map[string]string{
		"token": token,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})

}
