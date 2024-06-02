package entity

import "time"

type Transaction struct {
	Id      string
	RawDate time.Time
	Date    time.Time
	Amount  float32
	Label   string
}
