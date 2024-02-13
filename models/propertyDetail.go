package models

type PropertyDetail struct {
	Id                  uint `json:"productDetailsId"`
	ProductId           int  `json:"productId"`
	PropertyType        int  `json:"propertyType"`
	NumOfBedrooms       int  `json:"numOfBedrooms"`
	NumOfBathrooms      int  `json:"numOfBathrooms"`
	FloorArea           int  `json:"floorArea"`
	FurnishedStatus     bool `json:"furnishedStatus"`
	AdvancePayment      int  `json:"advancePayment"`
	SecurityDeposit     int  `json:"securityDeposit"`
	RentAmount          int  `json:"rentAmount"`
	Pool                bool `json:"pool"`
	Gym                 bool `json:"gym"`
	PropertyCode        int  `json:"propertyCode"`
	Generators          bool `json:"generators"`
	SeparateElectricity bool `json:"separateElectricity"`
}

func (productDetail *PropertyDetail) SetProductDetails(productId int, data ProductInput) {

	productDetail.ProductId = productId
	productDetail.PropertyType = data.PropertyType
	productDetail.NumOfBedrooms = data.NumOfBedrooms
	productDetail.NumOfBathrooms = data.NumOfBathrooms
	productDetail.FloorArea = data.FloorArea
	productDetail.FurnishedStatus = data.FurnishedStatus
	productDetail.AdvancePayment = data.AdvancePayment
	productDetail.SecurityDeposit = data.SecurityDeposit
	productDetail.RentAmount = data.RentAmount
	productDetail.PropertyCode = data.PropertyCode
	productDetail.Pool = data.Pool
	productDetail.Gym = data.Gym
	productDetail.Generators = data.Generators
	productDetail.SeparateElectricity = data.SeparateElectricity
}
