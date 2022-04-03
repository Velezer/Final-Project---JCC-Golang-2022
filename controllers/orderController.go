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
// @Success      200  {object}  models._Res{data=models.Order}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /orders [post]
func (h OrderController) CreateOrder(c *gin.Context) {
	var input OrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	userId := c.MustGet("user_id").(uint)

	savedOrder, err := services.All.OrderService.Save(userId, input.CartId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedOrder})
}

// GetOrders godoc
// @Summary      get Orders for user and merchant based on jwt
// @Description  get orders will check the role and show related orders
// @Tags         Order
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.Order}
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /orders [get]
func (h OrderController) GetOrders(c *gin.Context) {
	role := c.MustGet("user_role").(string)

	findingFunc := services.All.OrderService.FindAllByUserId
	if role == models.ROLE_MERCHANT {
		findingFunc = services.All.OrderService.FindAllByMerchantId
	}

	data, err := findingFunc(c.MustGet("user_id").(uint))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})

}

// GetOrder godoc
// @Summary      get Order based on jwt
// @Description  get order
// @Tags         Order
// @Param id path string true "order id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Order}
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /orders/{id} [get]
func (h OrderController) GetOrder(c *gin.Context) {
	data := c.MustGet("found").(*models.Order)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// DeleteOrder godoc
// @Summary      delete Order, user role must be MERCHANT
// @Description  delete order, only cancelled order can be deleted
// @Tags         Order
// @Param id path string true "order id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /orders/{id} [delete]
func (h OrderController) DeleteOrder(c *gin.Context) {
	orderId64 := c.MustGet("order_id").(uint64)
	orderId := uint(orderId64)
	found, err := services.All.OrderService.FindById(orderId)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
		return
	}
	if found.Status.Name != models.ORDER_CANCELLED {
		c.Error(errors.New("only CANCELLED order can be deleted")).SetMeta(httperror.NewMeta(http.StatusBadRequest))
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
// @Summary      update status Order, user role must be MERCHANT
// @Description  you can go right or left, but you can't revert. the starting point is UNPAID. ||| CANCELLED <- UNPAID -> PAID -> SHIPPING -> DELIVERED |||
// @Tags         Order
// @Param id path string true "order id"
// @Param        Body  body  updateStatusOrderInput  true  "the body to delete a Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Order}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
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
	if statusInput < statusNow && found.Status.Name != models.ORDER_UNPAID {
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

	orderId := c.MustGet("order_id").(uint64)
	data, err := services.All.OrderService.UpdateStatus(uint(orderId), input.Status)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
