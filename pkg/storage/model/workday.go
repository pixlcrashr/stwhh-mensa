// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage/datatypes"
	"gorm.io/gorm"
	"time"
)

type Workday struct {
	gorm.Model
	ID   datatypes.UUID `gorm:"primaryKey"`
	Date time.Time      `gorm:"uniqueIndex:workday_idx_date"`

	// relations
	GastronomyWorkdayDishes []GastronomyWorkdayDish
}
