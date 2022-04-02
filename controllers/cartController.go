package controllers

import (
	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
}

type CartInput struct {
	Name string `json:"name" binding:"required"`
}

// GetUserCart godoc
// @Summary      get cart, user role must be USER
// @Description  get user's cart with status is_checkout = false
// @Tags         Cart
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /carts [get]
func (h CartController) GetUserCart(c *gin.Context) {
	data, err := services.All.CartService.FindByuserId(c.MustGet("user_id").(uint))
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound).SetData("you don't have a cart"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// UpdateCart godoc
// @Summary      update cart, user role must be USER and must own the cart
// @Description  update cart name
// @Tags         Cart
// @Param id path string true "cart id"
// @Param        Body  body  CartInput  true  "the body to update a Cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /carts/{id} [put]
func (h CartController) UpdateCart(c *gin.Context) {
	var input CartInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	cartId := c.MustGet("cart_id").(uint)

	m := models.Cart{}
	m.Name = input.Name

	savedCart, err := services.All.CartService.Update(cartId, &m)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedCart})
}

// CreateCart godoc
// @Summary      Create Cart, user role must be USER
// @Description  create cart
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
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	m := models.Cart{}
	m.Name = input.Name
	m.UserId = c.MustGet("user_id").(uint)
	m.TotalPrice = 0

	savedCart, err := services.All.CartService.Save(&m)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedCart})
}

type CartItemInput struct {
	ProductId uint `json:"product_id" binding:"required"`
	Count     uint `json:"count" binding:"required"`
}

// AddCartItem godoc
// @Summary      add cart item, user role must be USER
// @Description  add cart item
// @Tags         Cart
// @Param id path string true "cart id"
// @Param        Body  body  CartItemInput  true  "the body to add a cart item"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /carts/{id}/items [post]
func (h CartController) AddCartItem(c *gin.Context) {
	var input CartItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}
	cartIdString := c.Param("id")
	cartId, err := strconv.ParseUint(cartIdString, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	item := models.CartItem{}
	item.CartId = uint(cartId)
	item.ProductId = input.ProductId
	item.Count = input.Count

	savedCart, err := services.All.CartService.AddCartItem(uint(cartId), item)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedCart})
}

// DeleteCartItem godoc
// @Summary      delete cart item, user role must be USER
// @Description  delete cart
// @Tags         Cart
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /carts/{cart_id}/items/{item_id} [delete]
func (h CartController) DeleteCartItem(c *gin.Context) {
	itemIdString := c.Param("item_id")
	itemId, err := strconv.ParseUint(itemIdString, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	savedCart, err := services.All.CartService.DeleteCartItem(uint(itemId))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedCart})
}
