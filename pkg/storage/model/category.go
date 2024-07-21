package model

import "gorm.io/gorm"
import "github.com/google/uuid"

type Category struct {
	gorm.Model
	ID      uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	STWHHID int       `gorm:"uniqueIndex:category_idx_stwhh_id"`
	Name    string    `gorm:"index:category_idx_name"`

	// relations
	DishCategories                  []DishCategory
	GastronomyWorkdayDishCategories []GastronomyWorkdayDishCategory
}
