package controllers

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "api health is good",
		"data":    nil,
		"success": true,
	})
}
