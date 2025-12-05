package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/inienam06/go-boilerplate/internal/core/jwt"
	"github.com/inienam06/go-boilerplate/internal/exception"
)

type BaseResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(exception.NewUnauthorizedException("Unauthorized"))
			c.Abort()
			return
		}

		// Extract the token (assuming "Bearer <token>")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse and validate the token
		token, err := jwt.ParseToken(tokenString)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("userID", token.UserID)
		c.Next()
	}
}
