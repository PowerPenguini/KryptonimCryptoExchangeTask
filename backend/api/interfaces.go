package api

import "github.com/shopspring/decimal"

type CurrencySvc interface {
	Rate(from, to string) (decimal.Decimal, error)
	Exchange(from, to string, amount decimal.Decimal) (decimal.Decimal, error)
}
