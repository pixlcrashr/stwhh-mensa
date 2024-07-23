// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage/datatypes"
	"gorm.io/gorm"
)

type GastronomyWorkdayDishPrice struct {
	gorm.Model
	ID                      datatypes.UUID `gorm:"primaryKey"`
	GastronomyWorkdayDishID datatypes.UUID `gorm:"uniqueIndex:gastronomy_workday_dish_price_idx_gastronomy_workday_dish_id_price_type"`
	PriceType               PriceType      `gorm:"uniqueIndex:gastronomy_workday_dish_price_idx_gastronomy_workday_dish_id_price_type"`
	Price                   int

	// relations
	GastronomyWorkdayDish GastronomyWorkdayDish
}
