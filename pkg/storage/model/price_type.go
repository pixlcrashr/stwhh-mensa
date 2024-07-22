// Package model
// Copyright © 2024 pixlcrashr (Vincent Heins)
package model

type PriceType int

const (
	GuestPriceType PriceType = iota
	StudentPriceType
	EmployeePriceType
)
