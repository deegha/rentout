package product

import (
	"fmt"
	"rentoutlkApi/databse"
	"rentoutlkApi/models"
	"rentoutlkApi/utils"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

/*
Add AddProduct
*/
func AddProduct(c *fiber.Ctx) error {

	// cookie := c.Cookies("jwt")
	token, err := utils.CheckAuth(c)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Authentication required",
			"data":    nil,
			"success": false,
		})
	}

	user := token

	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims, "claims")

	issuer := claims["iss"].(string)
	fmt.Println(issuer, "iss")

	// if user == nil {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Authentication required",
	// 		"data":    nil,
	// 		"success": false,
	// 	})
	// }

	var productInput models.ProductInput

	// parse inputs
	if err := c.BodyParser(&productInput); err != nil {
		fmt.Print(err.Error(), "Error in parsing request body")
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Error in parsing request body",
			"data":    nil,
			"success": false,
		})
	}

	isValid, e := productInput.IsValidProduct()

	if !isValid {
		return c.JSON(fiber.Map{
			"message": e,
			"data":    nil,
			"success": false,
		})
	}

	fmt.Println(productInput, "productInput")
	var product models.Product

	createdBy, _ := strconv.Atoi(issuer)
	product.CreatedBy = createdBy

	product.SetProduct(productInput)

	propertyCreateResult := databse.DB.Create(&product)

	if propertyCreateResult.Error != nil {
		fmt.Print(propertyCreateResult.Error, "Creating property")

		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": "Error in creating property",
			"data":    nil,
			"success": false,
		})
	}

	var wg sync.WaitGroup
	wg.Add(2)

	//Creating product details
	go func() {
		defer wg.Done()

		var propertyDetailInput = productInput
		var propertydetail models.PropertyDetail

		propertydetail.SetProductDetails(int(product.Id), propertyDetailInput)

		propertyDetailresult := databse.DB.Create(&propertydetail)

		if propertyDetailresult.Error != nil {
			fmt.Print(propertyDetailresult.Error, "Error in creating property details")
		}
	}()

	// adding productImages
	go addImages(productInput.Images, int(product.Id), &wg)

	wg.Wait()

	return c.JSON(fiber.Map{
		"message": "Successfully created",
		"data":    product,
		"success": true,
	})
}

func addImages(images []models.ProductImage, productId int, wg *sync.WaitGroup) {
	defer wg.Done()
	var imageArray []models.ImageDetail
	for _, img := range images {
		imageArray = append(imageArray, models.ImageDetail{Url: img.Url, ProductId: productId})
	}
	propetyImageResult := databse.DB.Create(&imageArray)

	if propetyImageResult.Error != nil {
		fmt.Print(propetyImageResult.Error, "Error in creating property images")
	}
}
