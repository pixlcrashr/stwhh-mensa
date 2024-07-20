package storage

import (
	"database/sql"
	model2 "github.com/pixlcrashr/stwhh-mensa/pkg/model"
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Storage struct {
	db *gorm.DB
}

func New(filepath string) (*Storage, error) {
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Dish{}); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) AddDish(dish model2.Dish) error {
	s.db.Create(&model.Dish{
		ID:   dish.ID,
		Name: dish.Name,
		PriceGuests: sql.NullInt32{
			Valid: dish.Prices.Guests.HasValue(),
			Int32: int32(dish.Prices.Guests.Value()),
		},
		PriceStudents: sql.NullInt32{
			Valid: dish.Prices.Students.HasValue(),
			Int32: int32(dish.Prices.Students.Value()),
		},
		PriceEmployees: sql.NullInt32{
			Valid: dish.Prices.Employees.HasValue(),
			Int32: int32(dish.Prices.Employees.Value()),
		},
	})
	return nil
}
