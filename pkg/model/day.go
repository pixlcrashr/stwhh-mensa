package model

import (
	"time"
)

type Day struct {
	Date         time.Time    `json:"date"`
	Gastronomies []Gastronomy `json:"gastronomies"`
}
