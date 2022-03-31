package middlewares

import (
	"hewantani/models"
	"hewantani/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id")
		u, err := services.All.UserService.FindById(userId.(uint))
		if err != nil {
			c.String(http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		if u.Role.Name != models.ROLE_USER {
			c.String(http.StatusForbidden, "role must be "+models.ROLE_USER)
			c.Abort()
			return
		}
		c.Next()
	}
}
