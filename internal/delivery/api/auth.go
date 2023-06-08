package api

import (
	"context"
	"github.com/namcchan/go-chatting/internal/domain"
	"github.com/namcchan/go-chatting/internal/middlewares"
	"github.com/namcchan/go-chatting/internal/usecase"
	"github.com/namcchan/go-chatting/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func AuthRegister(r *gin.RouterGroup) {
	router := r.Group("auth")

	router.POST("register", handleRegister)
	router.POST("login", handleLogin)
	router.POST("forgot-password", handleForgotPassword)
	router.POST("reset-password")
	router.GET("me", middlewares.CheckAuth, GetMe)
}

func handleRegister(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var registerPayload domain.RegisterPayload
	authUseCase := usecase.NewAuthUseCase(ctx)

	if err := c.BindJSON(&registerPayload); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if validationErr := validate.Struct(&registerPayload); validationErr != nil {
		c.JSON(http.StatusBadRequest, response.Error(validationErr.Error()))
		return
	}

	if err := authUseCase.Register(&registerPayload); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(""))
		return
	}

	c.JSON(http.StatusBadRequest, response.Success(nil))
}

func handleLogin(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	authUseCase := usecase.NewAuthUseCase(ctx)
	var payload domain.LoginPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "error": err.Error()})
		return
	}

	tokens, err := authUseCase.Login(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusBadRequest, response.Success(tokens))
}

func handleForgotPassword(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var payload domain.ForgotPasswordPayload
	authUseCase := usecase.NewAuthUseCase(ctx)

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if err := authUseCase.ForgotPassword(); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusBadRequest, response.Success(nil))
}

func GetMe(c *gin.Context) {
	c.JSON(http.StatusBadRequest, response.Success(c.MustGet("currentUser")))
}
