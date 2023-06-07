package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/namcchan/go-chatting/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateTokens(userId primitive.ObjectID, accessExpires int64, refreshExpires int64) (string, string, error) {

	fmt.Println(userId)
	accessClaims := jwt.MapClaims{
		"userId": userId.Hex(),
		"exp":    accessExpires,
	}

	refreshClaims := jwt.MapClaims{
		"userId": userId.Hex(),
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

func VerifyToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
