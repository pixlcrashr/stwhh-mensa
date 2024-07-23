// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage/datatypes"
	"gorm.io/gorm"
)

type DishCategory struct {
	gorm.Model
	ID         datatypes.UUID `gorm:"primaryKey"`
	DishID     datatypes.UUID `gorm:"uniqueIndex:dish_category_idx_dish_id_category_id"`
	CategoryID datatypes.UUID `gorm:"uniqueIndex:dish_category_idx_dish_id_category_id"`

	// relations
	Dish     Dish
	Category Category
}
