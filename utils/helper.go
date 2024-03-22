package utils

import (
	"github.com/golang-jwt/jwt"
)

const SecretKey = "secret"

func ValidateCookie(cookie string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}
