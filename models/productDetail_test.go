package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetProductDetails(t *testing.T) {

  data := map[string]string {
    "propertyType": "1",
    "numberOfBedrooms": "12",
    "numberOfBathrooms": "12",
    "floorArea": "1000",
    "furnishedStatus": "true",
    "advancePayment": "100000",
    "securityDeposite": "250000",
    "rentAmount": "2750000000",
    "pool": "true",
    "gym": "true",
    "propertyCode": "40222",
    "generators": "true",
    "seperateElectricity": "false",
}
  result := GetProductDetails(1,data )

  if(result.ProductId != 1) {
    t.Errorf("Failed to proccess the product id") 
  } else {
    t.Logf("Pass and proccessed the product id correctly")
  }
  if(result.PropertyType != 1) {
    t.Errorf("Failed to proccess the Property type correctly") 
  } else {
    t.Logf("Pass and proccessed the Property type correctly")
  }
  if(result.NumOfBedrooms!= 12) {
    t.Errorf("Failed to proccess theNumOfBedrooms correctly ") 
  } else {
    t.Logf("Pass and proccessed the theNumOfBedrooms correctly ")
  }

  values := reflect.ValueOf(result)
  types := values.Type()
  for i := 0; i < values.NumField(); i++ {
    fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
  }
}
