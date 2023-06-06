package api

import (
	"context"
	"github.com/namcchan/go-chatting/internal/domain"
	"github.com/namcchan/go-chatting/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// var userCollection *mongo.Collection = configs.GetCollection("users")
var validate = validator.New()

func AuthRegister(r *gin.RouterGroup) {
	router := r.Group("auth")

	router.POST("register", handleRegister)
	router.POST("login")
	router.POST("forgot-password")
	router.POST("reset-password")
	router.GET("me")
}

func handleRegister(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var registerPayload domain.RegisterPayload
	defer cancel()

	if err := c.BindJSON(&registerPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "error": err.Error()})
		return
	}

	if validationErr := validate.Struct(&registerPayload); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": validationErr.Error()})
		return
	}

	authUseCase := usecase.NewAuthUseCase(ctx)
	err := authUseCase.Register(&registerPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

//func handleLogin(c *gin.Context) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//}
