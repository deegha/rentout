package sevices

import (
	"errors"
	"fmt"
	"rentoutlkApi/databse"
	"strconv"
	"time"
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

type ListItem struct {
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
}

type Res struct {
	List            []ListItem
	TotalPages      int
	CurrentPage     int
	NumberOfRecords int
}

func FetchProperties(filter Filter) (*Res, error) {
	var page, _ = strconv.Atoi(filter.Page)

	fmt.Println("page", page)

	if page == 0 {
		page = 1
	}

	var resultsPerPage = 10
	var pageFirstResult = (page - 1) * resultsPerPage

	count, err := CountProperties(filter)
	if err != nil {
		return nil, err
	}

	var numberOfPages = (count + resultsPerPage - 1) / resultsPerPage

	// Using parameterized queries to prevent SQL injection
	query := `
		SELECT products.*, property_details.*
		FROM products
		INNER JOIN property_details ON property_details.id = products.id
		WHERE products.status = 1`

	var params []interface{}

	if filter.PropertyType != "" {
		query += " AND property_details.property_type = ?"
		params = append(params, filter.PropertyType)
	}
	if filter.NumOfBedrooms != "" {
		query += " AND property_details.num_of_bedrooms >= ?"
		params = append(params, filter.NumOfBedrooms)
	}
	if filter.NumOfBathrooms != "" {
		query += " AND property_details.num_of_bathrooms >= ?"
		params = append(params, filter.NumOfBathrooms)
	}
	if filter.FurnishedStatus != "" {
		query += " AND property_details.furnished_status = ?"
		params = append(params, filter.FurnishedStatus)
	}
	if filter.MaxRent != "" {
		query += " AND property_details.rent_amount <= ?"
		params = append(params, filter.MaxRent)
	}
	if filter.MinRent != "" {
		query += " AND property_details.rent_amount >= ?"
		params = append(params, filter.MinRent)
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", resultsPerPage, pageFirstResult)

	fmt.Println(query)

	var list []ListItem
	now := time.Now()
	result := databse.DB.Raw(query, params...).Scan(&list)

	fmt.Println(time.Since(now), "time taken to fetch properties from the database")

	var productCount int = len(list)

	if result.Error != nil {
		fmt.Print(result.Error)
		return nil, errors.New("error in fetching property details")
	}

	var res Res
	res.List = list
	res.TotalPages = numberOfPages
	res.CurrentPage = page
	res.NumberOfRecords = productCount

	return &res, nil
}

func CountProperties(filter Filter) (int, error) {
	query := `
		SELECT COUNT(*) as count
		FROM products
		INNER JOIN property_details ON property_details.id = products.id
		WHERE products.status = 1`

	var params []interface{}

	if filter.PropertyType != "" {
		query += " AND property_details.property_type = ?"
		params = append(params, filter.PropertyType)
	}
	if filter.NumOfBedrooms != "" {
		query += " AND property_details.num_of_bedrooms = ?"
		params = append(params, filter.NumOfBedrooms)
	}
	if filter.NumOfBathrooms != "" {
		query += " AND property_details.num_of_bathrooms = ?"
		params = append(params, filter.NumOfBathrooms)
	}
	if filter.FurnishedStatus != "" {
		query += " AND property_details.furnished_status = ?"
		params = append(params, filter.FurnishedStatus)
	}
	if filter.MaxRent != "" {
		query += " AND property_details.rent_amount <= ?"
		params = append(params, filter.MaxRent)
	}
	if filter.MinRent != "" {
		query += " AND property_details.rent_amount >= ?"
		params = append(params, filter.MinRent)
	}

	var countResult struct {
		Count int
	}

	result := databse.DB.Raw(query, params...).Scan(&countResult)

	if result.Error != nil {
		fmt.Print(result.Error)
		return 0, errors.New("error in fetching property count")
	}

	return countResult.Count, nil
}
