package transactions

import "errors"

var (
	ErrInvalidAccolcationRate = errors.New("invalid allocation rate, only number between 0 and 1 allowed")
)
