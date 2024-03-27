package user

import (
	"fmt"
	"rentoutlkApi/databse"
	"rentoutlkApi/models"
	"rentoutlkApi/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func User(c *fiber.Ctx) error {
	// cookie := c.Cookies("jwt")

	// token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(utils.SecretKey), nil
	// })
	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 		"data":    nil,
	// 		"success": false,
	// 	})
	// }

	token, err := utils.CheckAuth(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Authentication required",
			"data":    nil,
			"success": false,
		})
	}

	// claims := token.Claims.(*jwt.StandardClaims)
	claims := token.Claims.(jwt.MapClaims)
	var user models.User
	issuer := claims["iss"].(string)

	databse.DB.Where("id = ?", issuer).First(&user)

	fmt.Println(user, "user")

	if user.Id == 0 {
		return c.JSON(fiber.Map{
			"message": "No user found with this user id",
			"data":    nil,
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched user",
		"data":    user,
		"success": true,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var userInput models.UserInput

	if err := c.BodyParser(&userInput); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Error in parsing request body",
			"data":    nil,
			"success": false,
		})
	}

	IsValidUser, e := userInput.IsValidUser()

	if !IsValidUser {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": e,
			"data":    nil,
			"success": false,
		})
	}

	var user models.User
	user.SetPassword(string(userInput.Password))
	user.CreateConsumer(userInput)

	result := databse.DB.Create(&user)

	if result.Error != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Error in creating user",
			"data":    nil,
			"success": false,
		})
	}

	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"message": "Successfully created user",
		"data":    user,
		"success": true,
	})
}
