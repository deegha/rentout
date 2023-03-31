package controllers

import (
  "rentoutlkApi/models"
  "rentoutlkApi/databse"
  "github.com/gofiber/fiber/v2"
  "github.com/golang-jwt/jwt"
  "rentoutlkApi/utils"
)

func User(c *fiber.Ctx) error {
   cookie := c.Cookies("jwt")

   token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error) {
     return []byte(utils.SecretKey), nil})
	if err != nil {
	  c.Status(fiber.StatusUnauthorized)	
		  return c.JSON(fiber.Map{
	 "message": "unauthenticated",
		  })
	  }

  claims := token.Claims.(*jwt.StandardClaims)

  var user models.User

  databse.DB.Where("id = ?", claims.Issuer).First(&user)	

  return c.JSON(user)
}

