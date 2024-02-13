package models

type ImageDetail struct {
	Id        uint   `json:"id"`
	ProductId int    `json:"productId"`
	Url       string `json:"url"`
}
