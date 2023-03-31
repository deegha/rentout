package models

import (
	"strconv"
	"time"
)
type Product struct{
	Id uint`json:"id"`
	Title string`json:"title"`
	Description string`json:"description"`
	CreatedBy int`json:"createdBy"`
	ProductCategory int`json:"productCategory"`
	Status int`json:"status"`
	CreatedAt int64`json:"createdAt"`
	UpdatedAt int64`json:"updatedAt"`
}


func GetToBeUpdatedProduct(product Product, data map[string]string) Product{
	
	
  if(data["title"] != "") {
    product.Title = data["title"]
  }

  if(data["description"] != "") {
    product.Description = data["description"]
  }
 
  if(data["status"] != "") {
    status, _ := strconv.Atoi(data["status"])
    product.Status = status
  }

  product.UpdatedAt = time.Now().Unix()

  return product
}
