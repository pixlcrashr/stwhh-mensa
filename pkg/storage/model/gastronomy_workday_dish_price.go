// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GastronomyWorkdayDishPrice struct {
	gorm.Model
	ID                      uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	GastronomyWorkdayDishID uuid.UUID `gorm:"type:binary(128);uniqueIndex:gastronomy_workday_dish_price_idx_gastronomy_workday_dish_id_price_type"`
	PriceType               PriceType `gorm:"uniqueIndex:gastronomy_workday_dish_price_idx_gastronomy_workday_dish_id_price_type"`
	Price                   int

	// relations
	GastronomyWorkdayDish GastronomyWorkdayDish
}
