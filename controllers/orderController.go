package controllers

import (
	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

type OrderInput struct {
	UserId uint `json:"user_id" binding:"required"`
	CartId uint `json:"cart_id" binding:"required"`
}

// CreateOrder godoc
// @Summary      Create Order, user role must be USER
// @Description  registering a user from public access.
// @Tags         Order
// @Param        Body  body  OrderInput  true  "the body to create a Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /orders [post]
func (h OrderController) CreateOrder(c *gin.Context) {
	var input OrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	savedOrder, err := services.All.OrderService.Save(input.UserId, input.CartId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedOrder})
}

// GetOrders godoc
// @Summary      get Orders, anyone can use this
// @Description  registering a user from public access.
// @Tags         Order
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /orders [get]
func (h OrderController) GetOrders(c *gin.Context) {
	data, err := services.All.OrderService.FindAllByUserId(c.MustGet("user_id").(uint))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// DeleteOrder godoc
// @Summary      delete Order, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Order
// @Param        Body  body  OrderInput  true  "the body to delete a Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /orders/{id} [delete]
func (h OrderController) DeleteOrder(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	savedOrder, err := services.All.OrderService.Delete(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedOrder})
}

// CancelOrder godoc
// @Summary      cancel Order, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Order
// @Param        Body  body  OrderInput  true  "the body to delete a Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /orders/{id}/cancel [put]
func (h OrderController) CancelOrder(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	savedOrder, err := services.All.OrderService.UpdateStatus(uint(id), models.ORDER_CANCELLED)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedOrder})
}

// PayOrder godoc
// @Summary      cancel Order, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Order
// @Param        Body  body  OrderInput  true  "the body to delete a Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /orders/{id}/pay [put]
func (h OrderController) PayOrder(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	savedOrder, err := services.All.OrderService.UpdateStatus(uint(id), models.ORDER_COMPLETED)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedOrder})
}
