package productcontroller


import (
  "rentoutlkApi/models"
  "rentoutlkApi/utils"
  "strconv"
  "github.com/gofiber/fiber/v2"
)


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
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message" : "couldn't parse the data in the request body ",
      "data"    : nil,
    })
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


	product.CreateProduct()
  //Creating product details
	/*
  newProductDetails:= models.GetProductDetails(int(product.Id), data) 
  //adding productImatges
  images := strings.Split(data["images"], "+") 
  var imageArray []models.ImageDetail 
  for _, url := range images {
    imageArray = append(imageArray, models.ImageDetail{Url: url, ProductId: int(product.Id)})
  } 
  databse.DB.Create(&imageArray)	

  databse.DB.Create(&newProductDetailsl) */
  return c.JSON(fiber.Map{
    "message" : "Successfully created",
    "data"    : product,
  })
}
