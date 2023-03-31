package models

type ProductCategory struct{
	Id uint`json:"id"`
	CategoryId int`json:"categoryId"`
	CategoryName string`json:"categoryName"`
	CreatedAt int64`json:"createdAt"`
	UpdatedAt int64`json:"updatedAt"`

}
