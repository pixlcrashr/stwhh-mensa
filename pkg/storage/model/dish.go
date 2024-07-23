// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage/datatypes"
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID      datatypes.UUID `gorm:"primaryKey"`
	STWHHID int            `gorm:"uniqueIndex:dish_idx_stwhh_id"`
	Name    string         `gorm:"index:dish_idx_name"`

	// relations
	DishCategories []DishCategory
}
