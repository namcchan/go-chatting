package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/namcchan/go-chatting/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateTokens(userId primitive.ObjectID, accessExpires int64, refreshExpires int64) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    accessExpires,
	}

	refreshClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    refreshExpires,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenString, err := accessToken.SignedString([]byte(configs.GetEnv().AccessTokenSecret))
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(configs.GetEnv().RefreshTokenSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
