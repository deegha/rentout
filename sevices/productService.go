package sevices

import (
	"errors"
	"fmt"
	"rentoutlkApi/databse"
	"strconv"
)

type Filter struct {
		PropertyType    string
		NumOfBedrooms   string
		NumOfBathrooms  string
		FurnishedStatus string
		MaxRent         string
		MinRent         string
		Page            string
	}

type Res struct {
		ID              string `json:"id"`
		Title           string `json:"title"`
		Description     string `json:"description"`
		PropertyType    int    `json:"propertyType"`
		NumOfBedrooms   int    `json:"numOfBedrooms"`
		NumOfBathrooms  int    `json:"numOfBathrooms"`
		FurnishedStatus bool   `json:"furnishedStatus"`
		SecurityDeposit int    `json:"securityDeposit"`
		RentAmount      int    `json:"rentAmount"`
		Pool            bool   `json:"pool"`
		Gym             bool   `json:"gym"`		
	  Page            int    `json:"page"`
		NumberOfPage    int    `json:"numberOfPage"`
	}

func FetchProperties(filter Filter) ([]Res, error) {

	var page, _ = strconv.Atoi(filter.Page)

	if(page == 0) {
		page = 1 
	}
  var resultsPerPage int = 10
	var pageFirstResult = (page - 1)*resultsPerPage
  var productCount int64
	databse.DB.Table("products").Count(&productCount)

//  var numberOfPages = productCount/int64(resultsPerPage)
  
	query := `
		SELECT products.*, property_details.*
		FROM products
		INNER JOIN property_details ON property_details.id = products.id
		WHERE products.status = 1`

	if filter.PropertyType != "" {
		query += " AND property_details.property_type = '" + filter.PropertyType + "'"
	}
	if filter.NumOfBedrooms != "" {
		query += " AND property_details.num_of_bedrooms = " + filter.NumOfBedrooms
	}
	if filter.NumOfBathrooms != "" {
		query += " AND property_details.num_of_bathrooms = " + filter.NumOfBathrooms
	}
	if filter.FurnishedStatus != "" {
		query += " AND property_details.furnished_status = " + filter.FurnishedStatus
	}
	if filter.MaxRent != "" {
		query += " AND property_details.rent_amount <= " + filter.MaxRent
	}
	if filter.MinRent != "" {
		query += " AND property_details.rent_amount >= " + filter.MinRent
	}

  query += " limit "+fmt.Sprint(resultsPerPage)+" offset "+fmt.Sprint(pageFirstResult)
	
	fmt.Println(query)

	// Execute the query
	var res []Res
	result := databse.DB.Raw(query).Scan(&res)  

  if(result.Error != nil)	 {
		fmt.Print(result.Error)
	  return nil, errors.New("Error in fetching property details.")
	}

	return res, nil	
}
