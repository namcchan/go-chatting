package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/namcchan/go-chatting/configs"
	"github.com/namcchan/go-chatting/internal/repository"
	"github.com/namcchan/go-chatting/pkg/response"
	"github.com/namcchan/go-chatting/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func CheckAuth(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.JSON(http.StatusUnauthorized, response.Error("Unauthorized"))
		c.Abort()
		return
	}

	accessToken := strings.Replace(authorization, "Bearer ", "", -1)

	claims, err := utils.VerifyToken(accessToken, configs.GetEnv().AccessTokenSecret)
	if err != nil {
		c.Error(err)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(claims["userId"].(string))
	if err != nil {
		c.Error(err)
		return
	}

	filter := bson.M{"_id": objectID}
	projection := bson.M{"password": 0, "createdAt": 0, "updatedAt": 0}

	userRepo := repository.NewUserRepository(context.Background())
	user, err := userRepo.FindOne(filter, projection)

	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Error("Unauthorized"))
		c.Abort()
		return
	}

	c.Set("currentUser", user)

	c.Next()
}
