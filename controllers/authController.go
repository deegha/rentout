package controllers

import (
	"rentoutlkApi/databse"
	"rentoutlkApi/models"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
  "rentoutlkApi/utils"
)


/* 

Register function

*/
func Register(c *fiber.Ctx)	error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password:	password,
	}

	databse.DB.Create(&user)
	return c.JSON(user)
} 



/* 


Login function


*/
func Login(c *fiber.Ctx)	error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	databse.DB.Where("email = ?", data["email"]).First(&user)
	 
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect email",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
})

	token, err := claims.SignedString([]byte(utils.SecretKey))


	if err != nil {
	c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "couldnt sign user in",
		})
	}

	cookie := fiber.Cookie{

		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour*24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)


	return c.JSON(fiber.Map{
		"message": "Sucessfully logged in",
	})
}

/* 

Logout function

*/
func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
    "message": "Successfully logged out",
	})
}



