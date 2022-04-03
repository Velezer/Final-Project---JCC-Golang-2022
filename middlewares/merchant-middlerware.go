package middlewares

import (
	"errors"
	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MerchantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id").(uint)
		u, err := services.All.UserService.FindByIdJoinRole(userId)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
			c.Abort()
			return
		}

		if u.Role.Name != models.ROLE_MERCHANT {
			c.Error(errors.New("role must be " + models.ROLE_MERCHANT)).SetMeta(httperror.NewMeta(http.StatusForbidden))
			c.Abort()
			return
		}
		c.Next()
	}
}
