package domain

import "errors"

var (
	ErrInvalidAccolcationRate = errors.New("invalid allocation rate")
	ErrOverAllocation         = errors.New("too many allocations")
	ErrTransactionNotFound    = errors.New("no transaction found")
)
