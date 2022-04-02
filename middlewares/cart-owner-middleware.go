package middlewares

import (
	"hewantani/httperror"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CartOwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id").(uint)

		idString := c.Param("id")
		cartId, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
			c.Abort()
			return
		}

		found, err := services.All.CartService.FindById(uint(cartId))
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
			c.Abort()
			return
		}

		err = services.All.CartService.VerifyOwner(userId, found)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusForbidden))
			c.Abort()
			return
		}
		c.Set("found", found)
		c.Set("cart_id", uint(cartId))
		c.Next()
	}
}
