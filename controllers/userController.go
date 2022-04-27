package controllers

import (
	"errors"
	"net/http"

	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
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
// @Success      200  {object}  models._Res{data=map[string]string}
// @Success      400  {object}  models._Err
// @Router       /user [post]
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

	savedUser, err := services.All.UserService.Register(&u)
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
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res
// @Failure      400  {object}  models._Err
// @Failure      401  {object}  models._Err
// @Failure      500  {object}  models._Err
// @Router       /user/password [put]
func (a UserController) ChangePassword(c *gin.Context) {
	var input changePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	userId := c.MustGet("user_id")

	err := services.All.UserService.ChangePassword(userId.(uint), input.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed success"})

}

type updateUser struct {
	Email    string `json:"email" binding:"omitempty,email"` // gin binding is email but not required
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

// update user godoc
// @Summary      update user info
// @Description  update user info but can't change the role
// @Tags         User
// @Param        Body  body  updateUser  true  "the body to update a user"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=map[string]string}
// @Failure      400  {object}  models._Err
// @Failure      401  {object}  models._Err
// @Failure      404  {object}  models._Err
// @Failure      500  {object}  models._Err
// @Router       /user [put]
func (a UserController) UpdateUser(c *gin.Context) {
	var input updateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	if input == (updateUser{}) {
		c.Error(errors.New("input empty")).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	userId := c.MustGet("user_id").(uint)

	m := models.User{}
	m.Address = input.Address
	m.Username = input.Username
	m.Email = input.Email
	m.Password = input.Password

	user, err := services.All.UserService.Update(userId, &m)
	if err != nil {
		c.Error(err)
		return
	}

	data := map[string]string{}
	if user.Username != "" {
		data["username"] = user.Username
	}
	if user.Address != "" {
		data["address"] = user.Address
	}
	if user.Email != "" {
		data["email"] = user.Email
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated", "data": data})
}

// delete user godoc
// @Summary      delete user based on jwt
// @Description  delete user
// @Tags         User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res
// @Failure      401  {object}  models._Err
// @Failure      404  {object}  models._Err
// @Failure      500  {object}  models._Err
// @Router       /user [delete]
func (a UserController) DeleteUser(c *gin.Context) {
	userId := c.MustGet("user_id").(uint)
	err := services.All.UserService.Delete(userId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// Get User godoc
// @Summary      get user
// @Description  get user
// @Tags         User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=map[string]string{}}
// @Failure      401  {object}  models._Err
// @Failure      404  {object}  models._Err
// @Router       /user [get]
func (a UserController) GetUser(c *gin.Context) {
	userId := c.MustGet("user_id")

	user, err := services.All.UserService.FindByIdJoinRole(userId.(uint))
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
		return
	}

	data := map[string]string{
		"username": user.Username,
		"email":    user.Email,
		"address":  user.Address,
		"role":     user.Role.Name,
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed success", "data": data})

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
// @Success      200  {object}  models._Res{data=map[string]string{}}
// @Failure      400  {object}  models._Err
// @Router /user/login [post]
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
