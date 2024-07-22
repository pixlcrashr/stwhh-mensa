// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

type Dish struct {
	ID               int              `json:"id"`
	CategoryIDs      []int            `json:"category_ids"`
	Name             string           `json:"name"`
	Allergens        []string         `json:"allergens"`
	SymbolIDs        []int            `json:"symbol_ids"`
	Prices           Prices           `json:"prices"`
	EnvironmentScore EnvironmentScore `json:"environment_score"`
}
