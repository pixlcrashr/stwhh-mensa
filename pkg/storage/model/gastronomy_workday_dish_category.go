// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage/datatypes"
	"gorm.io/gorm"
)

type GastronomyWorkdayDishCategory struct {
	gorm.Model
	ID                      datatypes.UUID `gorm:"primaryKey"`
	GastronomyWorkdayDishID datatypes.UUID `gorm:"uniqueIndex:gastronomy_workday_dish_category_idx_gastronomy_workday_dish_id_category_id"`
	CategoryID              datatypes.UUID `gorm:"uniqueIndex:gastronomy_workday_dish_category_idx_gastronomy_workday_dish_id_category_id"`

	// relations
	GastronomyWorkdayDish GastronomyWorkdayDish
	Category              Category
}
