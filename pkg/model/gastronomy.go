// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

type Gastronomy struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Location   string     `json:"location"`
	Categories []Category `json:"categories"`
}
