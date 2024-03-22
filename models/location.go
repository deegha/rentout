package models

type Location struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
}
