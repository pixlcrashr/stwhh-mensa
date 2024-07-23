// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage/datatypes"
	"gorm.io/gorm"
)

type GastronomyWorkdayDish struct {
	gorm.Model
	ID           datatypes.UUID `gorm:"primaryKey"`
	GastronomyID datatypes.UUID `gorm:"uniqueIndex:gastronomy_workday_dish_idx_gastronomy_id_workday_id_dish_id_category_id"`
	WorkdayID    datatypes.UUID `gorm:"uniqueIndex:gastronomy_workday_dish_idx_gastronomy_id_workday_id_dish_id_category_id"`
	DishID       datatypes.UUID `gorm:"uniqueIndex:gastronomy_workday_dish_idx_gastronomy_id_workday_id_dish_id_category_id"`

	// relations
	Gastronomy Gastronomy
	Workday    Workday
	Dish       Dish
	Prices     []GastronomyWorkdayDishPrice
	Categories []GastronomyWorkdayDishCategory
}
