package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GastronomyWorkdayDish struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primaryKey;type:binary(128)"`
	GastronomyID uuid.UUID `gorm:"type:binary(128);uniqueIndex:gastronomy_workday_dish_idx_gastronomy_id_workday_id_dish_id_category_id"`
	WorkdayID    uuid.UUID `gorm:"type:binary(128);uniqueIndex:gastronomy_workday_dish_idx_gastronomy_id_workday_id_dish_id_category_id"`
	DishID       uuid.UUID `gorm:"type:binary(128);uniqueIndex:gastronomy_workday_dish_idx_gastronomy_id_workday_id_dish_id_category_id"`

	// relations
	Gastronomy Gastronomy
	Workday    Workday
	Dish       Dish
	Prices     []GastronomyWorkdayDishPrice
	Categories []GastronomyWorkdayDishCategory
}
