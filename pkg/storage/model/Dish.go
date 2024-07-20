package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID             int
	Name           string
	PriceGuests    sql.NullInt32
	PriceStudents  sql.NullInt32
	PriceEmployees sql.NullInt32
}
