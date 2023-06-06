package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Get the last error from the context
			lastError := c.Errors.Last()
			err := lastError.Err

			// Create an ErrorResponse object
			response := ErrorResponse{
				Message: "Internal Server Error",
				Error:   err.Error(),
			}

			// Set the response status code
			c.JSON(http.StatusInternalServerError, response)
		}
	}
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}
