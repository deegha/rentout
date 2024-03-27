package utils

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
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

func CheckAuth(c *fiber.Ctx) (*jwt.Token, error) {
	// Retrieve JWT token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		// No token provided
		return nil, errors.New("no token provided")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Return the key for verifying the token
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		// Return error if JWT token parsing fails or token is invalid
		return nil, errors.New("invalid token")
	}

	// Set the parsed token in the context for later use
	c.Locals("user", token)

	// Continue with next middleware or handler
	return token, nil
}
