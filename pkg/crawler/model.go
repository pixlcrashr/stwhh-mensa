// Package crawler
// Copyright © 2024 pixlcrashr (Vincent Heins)
package crawler

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/model"
	"time"
)

type Day struct {
	Date         time.Time          `json:"date"`
	Gastronomies []model.Gastronomy `json:"gastronomies"`
}
