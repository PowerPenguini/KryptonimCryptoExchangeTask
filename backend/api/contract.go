package api

import "github.com/shopspring/decimal"

type RateResponse struct {
	From string          `json:"from"`
	To   string          `json:"to"`
	Rate decimal.Decimal `json:"rate"`
}
