package model

import "github.com/pixlcrashr/stwhh-mensa/pkg/nullable"

type Prices struct {
	Students  nullable.Nullable[int] `json:"student"`
	Employees nullable.Nullable[int] `json:"employees"`
	Guests    nullable.Nullable[int] `json:"guests"`
}
