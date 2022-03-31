package middlewares

import (
	"hewantani/httperror"
	"hewantani/utils/jwttoken"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwttoken.TokenValid(c)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusUnauthorized))
			c.Abort()
			return
		}

		userId, err := jwttoken.ExtractTokenID(c)
		if err != nil {
			c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
