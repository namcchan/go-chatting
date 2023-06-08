package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/namcchan/go-chatting/pkg/response"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			statusCode := http.StatusInternalServerError
			if err.IsType(gin.ErrorTypePublic) {
				statusCode = err.Err.(*gin.Error).Meta.(int)
			}

			errorMessage := err.Err.(*gin.Error).Error()

			if statusCode != http.StatusInternalServerError {
				c.JSON(statusCode, response.Error(errorMessage))
			} else {
				c.JSON(http.StatusInternalServerError, response.Error("Internal Server Error"))
			}

			c.Abort()
		}
	}
}

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Error("API not found"))
	}
}
