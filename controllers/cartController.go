package controllers

import (
	"hewantani/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	Controller
}

type CartInput struct {
	Name string `json:"name" binding:"required"`
}

// CreateCart godoc
// @Summary      Create Cart, user role must be USER
// @Description  registering a user from public access.
// @Tags         Cart
// @Param        Body  body  CartInput  true  "the body to create a Cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /carts [post]
func (h CartController) CreateCart(c *gin.Context) {
	var input CartInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := models.Cart{}
	m.Name = input.Name
	m.UserId = c.MustGet("user_id").(uint)
	m.TotalPrice = 0

	savedCart, err := h.CartService.Save(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedCart})
}

type CartItemInput struct {
	ProductId uint `json:"product_id" binding:"required"`
	Count     uint `json:"count" binding:"required"`
}

// AddCartItem godoc
// @Summary      Create Cart, user role must be USER
// @Description  registering a user from public access.
// @Tags         Cart
// @Param        Body  body  CartInput  true  "the body to create a Cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /carts [post]
func (h CartController) AddCartItem(c *gin.Context) {
	var input CartItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cartIdString := c.Param("id")
	cartId, err := strconv.ParseUint(cartIdString, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item := models.CartItem{}
	item.CartId = uint(cartId)
	item.ProductId = input.ProductId
	item.Count = input.Count

	savedCart, err := h.CartService.AddCartItem(uint(cartId), item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedCart})
}
