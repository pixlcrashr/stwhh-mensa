// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GastronomyWorkdayDishCategory struct {
	gorm.Model
	ID                      uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	GastronomyWorkdayDishID uuid.UUID `gorm:"type:binary(128);uniqueIndex:gastronomy_workday_dish_category_idx_gastronomy_workday_dish_id_category_id"`
	CategoryID              uuid.UUID `gorm:"type:binary(128);uniqueIndex:gastronomy_workday_dish_category_idx_gastronomy_workday_dish_id_category_id"`

	// relations
	GastronomyWorkdayDish GastronomyWorkdayDish
	Category              Category
}
