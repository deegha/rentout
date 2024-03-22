package models

import (
	"fmt"
	"strconv"
	"time"
)

type Product struct {
	Id              uint   `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	CreatedBy       int    `json:"createdBy"`
	ProductCategory int    `json:"productCategory"`
	Status          int    `json:"status"`
	CreatedAt       int64  `json:"createdAt"`
	UpdatedAt       int64  `json:"updatedAt"`
}

type ProductImage struct {
	Url string `json:"url"`
}

type ProductInput struct {
	Id                  uint           `json:"id" validator:"required"`
	Title               string         `json:"title" validator:"required" binding:"required" `
	Description         string         `json:"description" validator:"required"`
	CreatedBy           int            `json:"createdBy"`
	ProductCategory     int            `json:"productCategory"`
	Status              int            `json:"status"`
	PropertyType        int            `json:"propertyType"`
	NumOfBedrooms       int            `json:"numOfBedrooms"`
	NumOfBathrooms      int            `json:"numOfBathrooms"`
	FloorArea           int            `json:"floorArea"`
	FurnishedStatus     bool           `json:"furnishedStatus"`
	AdvancePayment      int            `json:"advancePayment"`
	SecurityDeposit     int            `json:"securityDeposit"`
	RentAmount          int            `json:"rentAmount"`
	Pool                bool           `json:"pool"`
	Gym                 bool           `json:"gym"`
	PropertyCode        int            `json:"propertyCode"`
	Generators          bool           `json:"generators"`
	SeparateElectricity bool           `json:"separateElectricity"`
	CreatedAt           int64          `json:"createdAt"`
	UpdatedAt           int64          `json:"updatedAt"`
	Images              []ProductImage `json:"images"`
	ContactNumber       string         `json:"contactNumber"`
	ContactPerson       string         `json:"contactPerson"`
	LocationId          int            `json:"locationId"`
}

func GetToBeUpdatedProduct(product Product, data map[string]string) Product {
	if data["title"] != "" {
		product.Title = data["title"]
	}

	if data["description"] != "" {
		product.Description = data["description"]
	}

	if data["status"] != "" {
		status, _ := strconv.Atoi(data["status"])
		product.Status = status
	}

	product.UpdatedAt = time.Now().Unix()

	return product
}

func (product *Product) SetProduct(input ProductInput) {
	product.Title = input.Title
	product.Description = input.Description
	product.ProductCategory = input.ProductCategory
	product.Status = input.Status

	if input.CreatedBy != 0 {
		product.CreatedBy = input.CreatedBy
	}
}

func (productInput *ProductInput) IsValidProduct() (bool, string) {
	fmt.Println(productInput, "productInput")
	if productInput.Title == "" {
		return false, "title is required"
	}

	if productInput.ProductCategory == 0 {
		return false, "product Category is required"
	}

	if productInput.Status == 0 {
		return false, "status is required"
	}

	// if productInput.ContactNumber == "" {
	// 	return false, "contact number is required"
	// }

	// if productInput.ContactPerson == "" {
	// 	return false, "contact person is required"
	// }

	return true, ""
}
