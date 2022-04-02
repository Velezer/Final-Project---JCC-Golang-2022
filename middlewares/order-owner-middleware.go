package middlewares

import (
	"hewantani/httperror"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderOwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id").(uint)
		u, err := services.All.UserService.FindByIdJoinRole(userId)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusForbidden))
			c.Abort()
			return
		}

		idString := c.Param("id")
		orderId, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
			c.Abort()
			return
		}

		found, err := services.All.OrderService.FindById(uint(orderId))
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
			c.Abort()
			return
		}

		err = services.All.OrderService.VerifyOwner(userId, found)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusForbidden))
			c.Abort()
			return
		}
		c.Set("user_role", u.Role.Name)
		c.Set("found", found)
		c.Set("order_id", orderId)
		c.Next()
	}
}
