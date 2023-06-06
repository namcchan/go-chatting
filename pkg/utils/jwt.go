package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/namcchan/go-chatting/configs"
)

func GenerateTokens(userId string) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * 24 * 7).Unix(),
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
