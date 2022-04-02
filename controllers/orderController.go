package controllers

import (
	"errors"
	"hewantani/httperror"
	"hewantani/models"
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
// @Description  create order
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
// @Description  get orders
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
// @Description  delete order, only cancelled order can be deleted
// @Tags         Order
// @Param id path string true "order id"
// @Param        Body  body  OrderInput  true  "the body to delete a Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /orders/{id} [delete]
func (h OrderController) DeleteOrder(c *gin.Context) {
	orderId := c.MustGet("order_id").(uint)

	found, err := services.All.OrderService.FindById(orderId)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
		return
	}
	if found.Status.Name != models.ORDER_CANCELLED {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest).SetData(found.Status.Name + " order can't be deleted"))
		return
	}

	err = services.All.OrderService.Delete(orderId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success deleted order"})
}

type updateStatusOrderInput struct {
	Status string `json:"status" binding:"required,oneof=CANCELLED PAID SHIPPING DELIVERED"`
}

// UpdateStatusOrder godoc
// @Summary      pay || cancel Order, user role must be MERCHANT
// @Description  update order
// @Tags         Order
// @Param id path string true "order id"
// @Param        Body  body  updateStatusOrderInput  true  "the body to delete a Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /orders/{id} [put]
func (h OrderController) UpdateStatusOrder(c *gin.Context) {
	var input updateStatusOrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}
	userRole := c.MustGet("user_role").(string)
	found := c.MustGet("found").(*models.Order)

	// USER
	// UNPAID -> CANCELLED
	// UNPAID -> PAID
	// SHIPPING -> DELIVERED

	// MERCHANT
	// PAID -> SHIPPING
	getState := models.OrderStatus{}.GetState
	statusNow := getState(found.Status.Name)
	statusInput := getState(input.Status)
	if found.Status.Name == models.ORDER_CANCELLED {
		c.Error(errors.New("the status of cancelled order can't be updated")).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}
	if statusInput < statusNow {
		c.Error(errors.New("order status can't be reverted")).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}
	if statusInput > statusNow+1 {
		c.Error(errors.New("order status can't jump more than 2 level")).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}
	if input.Status == models.ORDER_SHIPPING && userRole == models.ROLE_USER {
		c.Error(errors.New("only merchant can ship the order")).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}
	if (input.Status == models.ORDER_PAID || input.Status == models.ORDER_DELIVERED) && userRole == models.ROLE_MERCHANT {
		c.Error(errors.New("only user can pay or mark the order as delivered")).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	orderId := c.MustGet("order_id").(uint)
	data, err := services.All.OrderService.UpdateStatus(orderId, input.Status)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
