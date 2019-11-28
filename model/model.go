package model

import "time"

type CBRResponse struct {
	Date         time.Time
	PreviousDate time.Time
	PreviousURL  string
	Timestamp    time.Time
	Currencies   map[string]Currency `json:"Valute"`
}

type Currency struct {
	ID       string
	NumCode  string
	CharCode string
	Nominal  int
	Name     string
	Value    float64
	Previous float64
}
