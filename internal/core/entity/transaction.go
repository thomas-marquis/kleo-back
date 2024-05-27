package entity

import "time"

type Transaction struct {
	Id      string
	RawDate time.Time
	Date    time.Time
	Amount  float64
	Label   string
}
