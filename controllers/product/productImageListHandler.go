package product

import (
	"fmt"
	"rentoutlkApi/databse"
	"rentoutlkApi/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ImageListHandler is a Fiber controller function to fetch images from the imageDetail model
func ListImages(c *fiber.Ctx) error {

	productId := c.Params("id")

	var images []models.ImageDetail

	now := time.Now()

	res := databse.DB.Where("product_id = ?", productId).Find(&images)

	//print the time taken to fetch the images since
	fmt.Println(time.Since(now), "time taken to fetch images")

	if res.Error != nil {
		fmt.Println(res.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch images",
			"data":    nil,
			"success": false,
		})
	}

	// Send the response with the correct HTTP code
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetched images",
		"data":    images,
		"success": true,
	})
}
