// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gastronomy struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	STWHHID  int       `gorm:"uniqueIndex:gastronomy_idx_stwhh_id"`
	Name     string    `gorm:"index:gastronomy_idx_name"`
	Location string    `gorm:"index:gastronomy_idx_location"`

	// relations
	GastronomyWorkdayDishes []GastronomyWorkdayDish
}
