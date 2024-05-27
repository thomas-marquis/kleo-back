package value

import (
	"time"
)

const (
	MaxItemsDefault = 100
)

type Filter struct {
	StartDate time.Time
	EndDate   time.Time
	MaxItems  int32
}

func NewFilter() Filter {
	return Filter{
		EndDate:  time.Now(),
		MaxItems: MaxItemsDefault,
	}
}
