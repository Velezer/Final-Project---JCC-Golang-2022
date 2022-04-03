package middlewares

import (
	"hewantani/httperror"
	"hewantani/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id").(uint)
		u, err := services.All.UserService.FindByIdJoinRole(userId)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
			c.Abort()
			return
		}

		c.Set("user_role", u.Role.Name)
		c.Next()
	}
}
