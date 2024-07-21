package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Workday struct {
	gorm.Model
	ID   uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	Date time.Time `gorm:"uniqueIndex:workday_idx_date"`

	// relations
	GastronomyWorkdayDishes []GastronomyWorkdayDish
}
