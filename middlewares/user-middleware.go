package middlewares

import (
	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("user_id")
		u, err := services.All.UserService.FindByIdJoinRole(userId.(uint))
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusForbidden))
			c.Abort()
			return
		}

		if u.Role.Name != models.ROLE_USER {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusForbidden).SetData("role must be " + models.ROLE_USER))
			c.Abort()
			return
		}
		c.Next()
	}
}
