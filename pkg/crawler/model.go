package crawler

import "github.com/pixlcrashr/stwhh-mensa/pkg/model"

type Result struct {
	Dishes       []model.Dish       `json:"dishes"`
	Categories   []model.Category   `json:"categories"`
	Gastronomies []model.Gastronomy `json:"gastronomies"`
	Allergens    []string           `json:"allergens"`
	Symbols      []model.Symbol     `json:"symbols"`
}
