package middlewares

import (
	"hewantani/models"
	"hewantani/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MerchantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id")
		u, err := services.All.UserService.FindByIdJoinRole(userId.(uint))
		if err != nil {
			c.String(http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		if u.Role.Name != models.ROLE_MERCHANT {
			c.String(http.StatusForbidden, "role must be "+models.ROLE_MERCHANT)
			c.Abort()
			return
		}
		c.Next()
	}
}
