package controllers

import (
	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartController struct {
}

type CartInput struct {
	Name string `json:"name" binding:"required"`
}

// GetUserCarts godoc
// @Summary      get carts, user role must be USER
// @Description  get user's carts with status is_checkout = false
// @Tags         Cart
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=interface{}}
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /carts [get]
func (h CartController) GetUserCarts(c *gin.Context) {
	found, err := services.All.CartService.FindAllByuserId(c.MustGet("user_id").(uint))
	if err != nil {
		c.Error(err)
		return
	}

	data := []map[string]interface{}{}
	for _, cart := range *found {
		data = append(data, map[string]interface{}{
			"id":          cart.ID,
			"name":        cart.Name,
			"is_checkout": cart.IsCheckout,
			"user_id":     cart.UserId,
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// GetUserCarts godoc
// @Summary      get carts, user role must be USER
// @Description  get user's cart with specific id
// @Tags         Cart
// @Param id path string true "cart id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Cart}
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /carts/{id} [get]
func (h CartController) GetUserCart(c *gin.Context) {
	found := c.MustGet("found").(*models.Cart)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": found})
}

// DeleteCart godoc
// @Summary      update cart, user role must be USER and must own the cart
// @Description  update cart name
// @Tags         Cart
// @Param id path string true "cart id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.Cart}
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /carts/{id} [delete]
func (h CartController) DeleteCart(c *gin.Context) {
	cartId := c.MustGet("cart_id").(uint)

	err := services.All.CartService.Delete(cartId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success deleted cart"})
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
// @Success      200  {object}  models._Res{data=[]models.Cart}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
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
// @Success      200  {object}  models._Res{data=[]models.Cart}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      500  {object}  models._Err
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
	Count     uint `json:"count" binding:"gte=0"`
}

// UpdateCartItem godoc
// @Summary      add item || update count || delete item, user role must be USER and must own the cart
// @Description  will insert if not exist (based on product_id), will update the count if exist, will delete if count is 0
// @Tags         Cart
// @Param id path string true "cart id"
// @Param        Body  body  CartItemInput  true  "the body to add a cart item"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /carts/{id}/items [put]
func (h CartController) UpdateCartItem(c *gin.Context) {
	var input CartItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	if _, err := services.All.ProductService.FindById(input.ProductId); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound).SetData("product id not found"))
		return
	}
	cartId := c.MustGet("cart_id").(uint)

	item := models.CartItem{}
	item.CartId = cartId
	item.ProductId = input.ProductId
	item.Count = input.Count

	err := services.All.CartService.UpdateCartItem(&item)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": map[string]interface{}{
		"cart_id":    cartId,
		"product_id": item.ProductId,
		"count":      item.Count,
	}})
}
