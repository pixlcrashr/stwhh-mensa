package model

type PriceType int

const (
	GuestPriceType PriceType = iota
	StudentPriceType
	EmployeePriceType
)
