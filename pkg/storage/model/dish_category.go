// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DishCategory struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	DishID     uuid.UUID `gorm:"type:binary(128);uniqueIndex:dish_category_idx_dish_id_category_id"`
	CategoryID uuid.UUID `gorm:"type:binary(128);uniqueIndex:dish_category_idx_dish_id_category_id"`

	// relations
	Dish     Dish
	Category Category
}
