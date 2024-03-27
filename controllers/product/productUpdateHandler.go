package product

import (
	"fmt"
	"rentoutlkApi/databse"
	"rentoutlkApi/models"
	"rentoutlkApi/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

/*
update Product
*/
func UpdateProduct(c *fiber.Ctx) error {
	var data map[string]string
	// cookie := c.Cookies("jwt")
	// _, err := utils.ValidateCookie(cookie)

	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)

	// 	return c.JSON(fiber.Map{
	// 		"message": "Authentication reuired",
	// 		"data":    nil,
	// 	})
	// }

	// cookie := c.Cookies("jwt")
	_, err := utils.CheckAuth(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Authentication required",
			"data":    nil,
			"success": false,
		})
	}

	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Somthing wrong in the payload",
			"data":    nil,
		})
	}

	productId := c.Params("id")

	var product models.Product
	databse.DB.Where("id = ?", productId).First(&product)

	if data["title"] != "" {
		product.Title = data["title"]
	}

	if data["description"] != "" {
		product.Description = data["description"]
	}

	status, _ := strconv.Atoi(data["status"])
	if data["status"] != "" {
		product.Status = status
	}

	product.UpdatedAt = time.Now().Unix()

	created := databse.DB.Save(&product)
	if created.Error != nil {
		fmt.Println(created.Error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Coulndt update the property",
			"data":    nil,
		})
	}
	var pId, _ = strconv.Atoi(strings.TrimSpace(productId))
	var details models.PropertyDetail
	databse.DB.Where("product_id = ?", pId).First(&details)

	var propertyType, _ = strconv.Atoi(data["propertyType"])
	var numberOfBedrooms, _ = strconv.Atoi(data["numberOfBedrooms"])
	var numberOfBathrooms, _ = strconv.Atoi(data["numberOfBathrooms"])
	var floorArea, _ = strconv.Atoi(data["floorArea"])
	var furnishedStatus, _ = strconv.ParseBool(data["furnishedStatus"])
	var advancePayment, _ = strconv.Atoi(data["advancePayment"])
	var securityDeposite, _ = strconv.Atoi(data["securityDeposite"])
	var rentAmount, _ = strconv.Atoi(data["rentAmount"])
	var propertyCode, _ = strconv.Atoi(data["propertyCode"])
	var pool, _ = strconv.ParseBool(data["pool"])
	var gym, _ = strconv.ParseBool(data["gym"])
	var generators, _ = strconv.ParseBool(data["generators"])
	var seperateElectricity, _ = strconv.ParseBool(data["seperateElectricity"])

	if propertyType != 0 {
		details.PropertyType = propertyType
	}
	if numberOfBedrooms != 0 {
		details.NumOfBedrooms = numberOfBedrooms
	}
	if numberOfBathrooms != 0 {
		details.NumOfBathrooms = numberOfBathrooms
	}
	if floorArea != 0 {
		details.FloorArea = floorArea
	}
	if furnishedStatus != details.FurnishedStatus {
		details.FurnishedStatus = furnishedStatus
	}
	if advancePayment != 0 {
		details.AdvancePayment = advancePayment
	}
	if securityDeposite != 0 {
		details.SecurityDeposit = securityDeposite
	}
	if rentAmount != 0 {
		details.RentAmount = rentAmount
	}
	if propertyCode != 0 {
		details.PropertyCode = propertyCode
	}
	if pool != details.Pool {
		details.Pool = pool
	}
	if gym {
		details.Gym = gym
	}
	if generators != details.Generators {
		details.Generators = generators
	}
	if seperateElectricity != details.SeparateElectricity {
		details.SeparateElectricity = seperateElectricity
	}

	fmt.Println(details.Id)

	// newProductDetails:= models.GetProductDetails(pId, data)
	result := databse.DB.Updates(&details)

	if result.Error != nil {
		fmt.Println(created.Error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Coulndn't update the property details",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated",
		"data":    product,
	})
}
