// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID      uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	STWHHID int       `gorm:"uniqueIndex:dish_idx_stwhh_id"`
	Name    string    `gorm:"index:dish_idx_name"`

	// relations
	DishCategories []DishCategory
}
