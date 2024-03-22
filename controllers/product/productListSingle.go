package product

import (
	"rentoutlkApi/databse"

	"github.com/gofiber/fiber/v2"
)

/*
List a Products
*/
func ListAProducts(c *fiber.Ctx) error {
	type Res struct {
		Id                  string `json:"id"`
		Title               string `json:"title"`
		Description         string `json:"description"`
		CreatedBy           int    `json:"createdBy"`
		Status              int    `json:"status"`
		CreatedAt           int64  `json:"createdAt"`
		PropertyType        int    `json:"propertyType"`
		NumOfBedrooms       int    `json:"numOfBedrooms"`
		NumOfBathrooms      int    `json:"numOfBathrooms"`
		FloorArea           int    `json:"floorArea"`
		FurnishedStatus     bool   `json:"furnishedStatus"`
		AdvancePayment      int    `json:"advancePayment"`
		SecurityDeposit     int    `json:"securityDeposit"`
		RentAmount          int    `json:"rentAmount"`
		Pool                bool   `json:"pool"`
		Gym                 bool   `json:"gym"`
		PropertyCode        int    `json:"propertyCode"`
		Generators          bool   `json:"generators"`
		SeperateElectricity bool   `json:"seperateElectricity"`
		ContactNumber       string `json:"contactNumber"`
		ContactPerson       string `json:"contactPerson"`
		LocationId          int    `json:"locationId"`
	}

	productId := c.Params("id")
	var res Res
	result := databse.DB.Raw("select * from products inner join property_details on property_details.id = products.id where products.id=?", productId).Scan(&res)

	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Couln'd fetch data",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"message": "List of products",
		"data":    res,
	})
}
