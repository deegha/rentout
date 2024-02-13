package product

import (
	"fmt"
	"rentoutlkApi/sevices"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ListProducts(c *fiber.Ctx) error {
	// Extracting query parameters to create a filter
	filter := sevices.Filter{
		PropertyType:    c.Query("propertyType"),
		NumOfBedrooms:   c.Query("beds"),
		NumOfBathrooms:  c.Query("baths"),
		FurnishedStatus: c.Query("furnished"),
		MaxRent:         c.Query("maxRent"),
		MinRent:         c.Query("minRent"),
		Page:            c.Query("page"),
	}

	// Call FetchProperties function
	now := time.Now()
	properties, err := sevices.FetchProperties(filter)
	//print the time taken to fetch the properties from the service since
	fmt.Println(time.Since(now), "time taken to fetch properties from the service since")

	// Check for errors
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println(err)

		// Return an appropriate HTTP status and response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching properties",
			"data":    nil,
		})
	}

	// Check if the properties list is nil or empty
	if properties == nil || len(properties.List) == 0 {
		// Return a specific message for no data
		return c.JSON(fiber.Map{
			"message": "No properties found",
			"data":    nil,
		})
	}

	// Return the list of properties
	return c.JSON(fiber.Map{
		"message": "List of properties",
		"data":    properties,
	})
}
