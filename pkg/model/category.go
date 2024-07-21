package model

type Category struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Dishes []Dish `json:"dishes"`
}
