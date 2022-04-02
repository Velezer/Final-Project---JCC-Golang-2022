package middlewares

import (
	"hewantani/httperror"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StoreOwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id").(uint)

		idString := c.Param("id")
		storeId, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
			c.Abort()
			return
		}

		found, err := services.All.StoreService.FindById(uint(storeId))
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
			c.Abort()
			return
		}

		err = services.All.StoreService.VerifyOwner(userId, found)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusForbidden))
			c.Abort()
			return
		}
		c.Set("found", found)
		c.Set("store_id", uint(storeId))
		c.Next()
	}
}
