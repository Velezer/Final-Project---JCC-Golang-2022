package controllers

import (
	"hewantani/services"
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedOrder, err := services.All.OrderService.Save(input.UserId, input.CartId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedOrder})
}
