package middleware

import (
	"github.com/gin-gonic/gin"
	"vtmtea.com/fiction/handler"
	"vtmtea.com/fiction/pkg/errno"
	"vtmtea.com/fiction/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
