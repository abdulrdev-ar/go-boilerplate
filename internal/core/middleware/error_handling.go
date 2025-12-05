package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inienam06/go-boilerplate/internal/exception"
	"github.com/inienam06/go-boilerplate/internal/response"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		// cek jenis error
		ex, ok := err.Err.(*exception.HttpException)
		if ok {
			c.JSON(ex.StatusCode, response.BaseResponse{
				Status:  false,
				Message: ex.Message,
				Data:    nil,
			})
			return
		}

		// default internal server error
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
}
