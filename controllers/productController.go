package controllers

import (
	"fmt"
	"rentoutlkApi/sevices"
	"rentoutlkApi/databse"
	"rentoutlkApi/models"
	"rentoutlkApi/utils"
	"strconv"
	"time"
	"strings"
	"github.com/gofiber/fiber/v2"
)

/*

Add AddProduct

*/


func AddProduct(c *fiber.Ctx) error {
	var data map[string]string
  cookie := c.Cookies("jwt")
  claims, err := utils.ValidateCookie(cookie) 

  if err != nil {
    return c.JSON(fiber.Map{
      "message" : "Authentication reuired",
      "data"    : nil,
    })
  }

	if err := c.BodyParser(&data); err != nil {
		return err
	}

  createdBy, _ := strconv.Atoi(claims.Issuer)
  productCategory, _ := strconv.Atoi(data["productCategory"])

  // Creating the product
  product := models.Product{
    Title           : data["title"],
    Description     : data["description"],
    CreatedBy       : createdBy,
    ProductCategory : productCategory,
  }

  databse.DB.Create(&product)

	//Creating product details
  newProductDetails:= models.GetProductDetails(int(product.Id), data) 
	databse.DB.Create(&newProductDetails)
		return c.JSON(fiber.Map{
    "message" : "Successfully created",
    "data"    : product,
  })
}

/*

List a Products

*/
func ListAProducts(c *fiber.Ctx) error {

  type Res struct {
		Id                  string`json:"id"`
		Title               string`json:"title"`
		Description         string`json:"description"`
		CreatedBy           int`json:"createdBy"`
		Status              int`json:"status"`
		CreatedAt           int64`json:"createdAt"`
		PropertyType        int`json:"propertyType"`
		NumOfBedrooms       int`json:"numOfBedrooms"`
		NumOfBathrooms      int`json:"numOfBathrooms"`
		FloorArea           int`json:"floorArea"`
		FurnishedStatus     bool`json:"furnishedStatus"`
		AdvancePayment      int`json:"advancePayment"`
		SecurityDeposit     int`json:"securityDeposit"`
		RentAmount          int`json:"rentAmount"`
		Pool                bool`json:"pool"`
		Gym                 bool`json:"gym"`
		PropertyCode        int`json:"propertyCode"`
		Generators          bool`json:"generators"`
		SeperateElectricity bool`json:"seperateElectricity"`
	}

	
  productId := c.Params("id")
	var res Res
  result :=	databse.DB.Raw("select * from products inner join property_details on property_details.id = products.id where products.id=?", productId).Scan(&res)
	
  if(result.Error != nil) {
		c.Status(fiber.StatusInternalServerError)
    return c.JSON(fiber.Map{
      "message": "Couln'd fetch data",
      "data"   : nil,
    })
  }

  return c.JSON(fiber.Map{
    "message": "List of products",
		"data"   : res,
  })
}
/*

update Product

*/
func UpdateProduct(c *fiber.Ctx) error {
	var data map[string]string
  cookie := c.Cookies("jwt")
  _, err := utils.ValidateCookie(cookie) 

  if err != nil {
		c.Status(fiber.StatusUnauthorized)

    return c.JSON(fiber.Map{
      "message" : "Authentication reuired",
      "data"    : nil,
    })
  }

	if err := c.BodyParser(&data); err != nil {
   fmt.Println(err) 
		c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message": "Somthing wrong in the payload",
      "data": nil,
    })

	}

  productId := c.Params("id")

  var product models.Product;
  databse.DB.Where("id = ?", productId ).First(&product)	

	if(data["title"] != "")  {
		product.Title = data["title"]
	}

	if(data["description"] != "") {
		product.Description = data["description"]
	}
	
	status, _:= strconv.Atoi(data["status"])
	if(data["status"] != "") {
		product.Status = status
	}

  product.UpdatedAt = time.Now().Unix()

  created := databse.DB.Save(&product)
  if(created.Error != nil) {
    fmt.Println(created.Error) 
		c.Status(fiber.StatusInternalServerError)
    return c.JSON(fiber.Map{
      "message": "Coulndt update the property",
      "data": nil,
    })
  }
  var pId, _ = strconv.Atoi(strings.TrimSpace(productId))
	var details models.PropertyDetail 
	databse.DB.Where("product_id = ?", pId).First(&details)	


  var propertyType, _        = strconv.Atoi(data["propertyType"]) 
  var numberOfBedrooms, _    = strconv.Atoi(data["numberOfBedrooms"]) 
  var numberOfBathrooms, _   = strconv.Atoi(data["numberOfBathrooms"]) 
  var floorArea, _           = strconv.Atoi(data["floorArea"]) 
  var furnishedStatus, _     = strconv.ParseBool(data["furnishedStatus"]) 
  var advancePayment, _      = strconv.Atoi(data["advancePayment"]) 
  var securityDeposite, _    = strconv.Atoi(data["securityDeposite"]) 
  var rentAmount, _          = strconv.Atoi(data["rentAmount"]) 
  var propertyCode, _        = strconv.Atoi(data["propertyCode"])
  var pool, _                = strconv.ParseBool(data["pool"])
  var gym, _                 = strconv.ParseBool(data["gym"])
  var generators, _          = strconv.ParseBool(data["generators"])
  var seperateElectricity, _ = strconv.ParseBool(data["seperateElectricity"])

	if(propertyType != 0 ) {
	  details.PropertyType = propertyType
	}
	if(numberOfBedrooms != 0) {
			details.NumOfBedrooms= numberOfBedrooms
	}
	if(numberOfBathrooms != 0 ) {
	  details.NumOfBathrooms = numberOfBathrooms
	}
  if(floorArea != 0 ) {
	  details.FloorArea= floorArea
	}
	if(furnishedStatus != details.FurnishedStatus) {
			details.FurnishedStatus= furnishedStatus 
	}
	if(advancePayment != 0) {
	  details.AdvancePayment = advancePayment
	}
	if(securityDeposite != 0) {
	  details.SecurityDeposit = securityDeposite 
	}
	if(rentAmount != 0) {
	  details.RentAmount = rentAmount 
	}
  if(propertyCode !=  0) {
	  details.PropertyCode = propertyCode 
	}
	if(pool != details.Pool) {
			details.Pool = pool 
	}
	if(gym) {
	  details.Gym = gym 
	}
	if(generators != details.Generators) {
			details.Generators = generators 
	}
	if(seperateElectricity != details.SeperateElectricity) {
	  details.SeperateElectricity = seperateElectricity 
	}


fmt.Println(details.Id)

 // newProductDetails:= models.GetProductDetails(pId, data) 
  result := databse.DB.Updates(&details) 

  if(result.Error != nil) {
    fmt.Println(created.Error) 
		c.Status(fiber.StatusInternalServerError)
    return c.JSON(fiber.Map{
      "message": "Coulndn't update the property details",
      "data": nil,
    })
  }

	return c.JSON(fiber.Map{
			"message": "Successfully updated",
    "data": product,
  })
}

/*

List Product

*/
func ListProducts(c *fiber.Ctx) error {
  
/**
  adding filters
*/

fmt.Println(c.Query("productType"))
	filter := sevices.Filter{
		PropertyType:    c.Query("propertyType"),
		NumOfBedrooms:   c.Query("beds"),
		NumOfBathrooms:  c.Query("baths"),
		FurnishedStatus: c.Query("furnished"),
		MaxRent:         c.Query("maxRent"),
		MinRent:         c.Query("minRent"),
		Page:            c.Query("page"),
	}
  
	properties, err := sevices.FetchProperties(filter)
	
  if(err != nil) {
		c.Status(fiber.StatusInternalServerError)

		//loggin the error 
		fmt.Println(err)
    return c.JSON(fiber.Map{
      "message": err,
      "data": nil,
    })
  }

  return c.JSON(fiber.Map{
    "message": "List of products",
		"data": properties,
  })
}
