package models
import (
  "strconv"
)

type PropertyDetail struct {
  Id uint`json:"productDetailsId"`
  ProductId int`js:"productId"`
  PropertyType int`json:"propertyType"`
  NumOfBedrooms int`json:"numOfBedrooms"`
  NumOfBathrooms int`json:"numOfBathrooms"`
  FloorArea int`json:"floorArea"`
  FurnishedStatus bool`json:"furnishedStatus"`
  AdvancePayment int`json:"advancePayment"`
  SecurityDeposit int`json:"securityDeposit"`
  RentAmount int`json:"rentAmount"`
  Pool bool`json:"pool"`
  Gym bool`json:"gym"`
  PropertyCode int`json:"propertyCode"`
  Generators bool`json:"generators"`
  SeperateElectricity bool`json:"seperateElectricity"`
}


func (productDetail  *PropertyDetail) GetProductDetails(productId int, data map[string]string){ 

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

  //Creating the product Details
    productDetail.ProductId = productId
    productDetail.PropertyType = propertyType
    productDetail.NumOfBedrooms = numberOfBedrooms
    productDetail.NumOfBathrooms = numberOfBathrooms
    productDetail.FloorArea = floorArea
    productDetail.FurnishedStatus = furnishedStatus
    productDetail.AdvancePayment = advancePayment
    productDetail.SecurityDeposit = securityDeposite
    productDetail.RentAmount = rentAmount
    productDetail.PropertyCode = propertyCode
    productDetail.Pool = pool
    productDetail.Gym = gym
    productDetail.Generators = generators
    productDetail.SeperateElectricity = seperateElectricity

}
