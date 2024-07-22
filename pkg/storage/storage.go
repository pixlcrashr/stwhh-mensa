// Package storage
// Copyright © 2024 pixlcrashr (Vincent Heins)
package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/pixlcrashr/stwhh-mensa/pkg/crawler"
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

	if err := db.AutoMigrate(
		&model.Category{},
		&model.Dish{},
		&model.DishCategory{},
		&model.Gastronomy{},
		&model.GastronomyWorkdayDish{},
		&model.GastronomyWorkdayDishCategory{},
		&model.GastronomyWorkdayDishPrice{},
		&model.Workday{},
	); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) AddDay(ctx context.Context, day crawler.Day) (err error) {
	var workday model.Workday

	if err = s.db.
		WithContext(ctx).
		Where("date = ?", day.Date).
		Attrs(model.Workday{
			ID:   uuid.New(),
			Date: day.Date,
		}).
		FirstOrCreate(&workday).
		Error; err != nil {
		return err
	}

	for _, gastronomy := range day.Gastronomies {
		var gastronomyModel model.Gastronomy
		if err := s.db.
			WithContext(ctx).
			Attrs(&model.Gastronomy{
				ID:       uuid.New(),
				STWHHID:  gastronomy.ID,
				Name:     gastronomy.Name,
				Location: gastronomy.Location,
			}).
			FirstOrCreate(&gastronomyModel, model.Gastronomy{
				STWHHID: gastronomy.ID,
			}).
			Error; err != nil {
			return err
		}

		for _, category := range gastronomy.Categories {
			var categoryModel model.Category
			if err := s.db.
				WithContext(ctx).
				Attrs(&model.Category{
					ID:      uuid.New(),
					STWHHID: category.ID,
					Name:    category.Name,
				}).
				FirstOrCreate(&categoryModel, model.Category{
					STWHHID: category.ID,
				}).
				Error; err != nil {
				return err
			}

			for _, dish := range category.Dishes {
				var dishModel model.Dish
				if err := s.db.
					WithContext(ctx).
					Attrs(&model.Dish{
						ID:      uuid.New(),
						STWHHID: dish.ID,
						Name:    dish.Name,
					}).
					FirstOrCreate(&dishModel, model.Dish{
						STWHHID: dish.ID,
					}).
					Error; err != nil {
					return err
				}

				var gastronomyWorkdayDishModel model.GastronomyWorkdayDish
				if err := s.db.
					WithContext(ctx).
					Attrs(&model.GastronomyWorkdayDish{
						ID:           uuid.New(),
						DishID:       dishModel.ID,
						GastronomyID: gastronomyModel.ID,
						WorkdayID:    workday.ID,
					}).
					FirstOrCreate(&gastronomyWorkdayDishModel, model.GastronomyWorkdayDish{
						DishID:       dishModel.ID,
						GastronomyID: gastronomyModel.ID,
						WorkdayID:    workday.ID,
					}).
					Error; err != nil {
					return err
				}

				if err := s.db.
					WithContext(ctx).
					Attrs(&model.GastronomyWorkdayDishCategory{
						ID:                      uuid.New(),
						GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
						CategoryID:              categoryModel.ID,
					}).
					FirstOrCreate(&model.GastronomyWorkdayDishCategory{}, model.GastronomyWorkdayDishCategory{
						GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
						CategoryID:              categoryModel.ID,
					}).
					Error; err != nil {
					return err
				}

				var prices = dish.Prices
				if prices.Guests.HasValue() {
					if err := s.db.
						WithContext(ctx).
						Attrs(&model.GastronomyWorkdayDishPrice{
							ID:                      uuid.New(),
							GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
							PriceType:               model.GuestPriceType,
							Price:                   dish.Prices.Guests.Value(),
						}).
						FirstOrCreate(&model.GastronomyWorkdayDishPrice{}, model.GastronomyWorkdayDishPrice{
							GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
							PriceType:               model.GuestPriceType,
						}).
						Error; err != nil {
						return err
					}
				}

				if prices.Students.HasValue() {
					if err := s.db.
						WithContext(ctx).
						Attrs(&model.GastronomyWorkdayDishPrice{
							ID:                      uuid.New(),
							GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
							PriceType:               model.StudentPriceType,
							Price:                   dish.Prices.Students.Value(),
						}).
						FirstOrCreate(&model.GastronomyWorkdayDishPrice{}, model.GastronomyWorkdayDishPrice{
							GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
							PriceType:               model.StudentPriceType,
						}).
						Error; err != nil {
						return err
					}
				}

				if prices.Employees.HasValue() {
					if err := s.db.
						WithContext(ctx).
						Attrs(&model.GastronomyWorkdayDishPrice{
							ID:                      uuid.New(),
							GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
							PriceType:               model.EmployeePriceType,
							Price:                   dish.Prices.Employees.Value(),
						}).
						FirstOrCreate(&model.GastronomyWorkdayDishPrice{}, model.GastronomyWorkdayDishPrice{
							GastronomyWorkdayDishID: gastronomyWorkdayDishModel.ID,
							PriceType:               model.EmployeePriceType,
						}).
						Error; err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
