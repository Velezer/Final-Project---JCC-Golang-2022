package middlewares

import (
	"hewantani/httperror"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductOwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id").(uint)

		idString := c.Param("id")
		productId, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
			c.Abort()
			return
		}

		found, err := services.All.ProductService.FindById(uint(productId))
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
			c.Abort()
			return
		}

		err = services.All.ProductService.VerifyOwner(userId, found)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusForbidden))
			c.Abort()
			return
		}
		c.Set("found", found)
		c.Set("product_id", uint(productId))
		c.Next()
	}
}
