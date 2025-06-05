package helpers

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID uint, role string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iat":  time.Now().Unix(),
		"sub":  userID,
		"role": role,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateJWT_RS256(userID uint, role string, privateKey *rsa.PrivateKey, kid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iat":  time.Now().Unix(),
		"sub":  userID,
		"role": role,
	})

	token.Header["kid"] = kid

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
