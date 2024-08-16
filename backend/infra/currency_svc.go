package infra

import (
	"errors"

	"github.com/shopspring/decimal"
)

type currency struct {
	Symbol     string
	Precision  int32
	RateToBase decimal.Decimal
}

const maxPrecision = 18 // I know how it looks like, let's talk ;)

var currencies = []currency{
	{Symbol: "USD", Precision: 2, RateToBase: decimal.NewFromFloat(1.0)},
	{Symbol: "BEER", Precision: 18, RateToBase: decimal.NewFromFloat(0.00002461)},
	{Symbol: "FLOKI", Precision: 18, RateToBase: decimal.NewFromFloat(0.0001428)},
	{Symbol: "GATE", Precision: 18, RateToBase: decimal.NewFromFloat(6.87)},
	{Symbol: "USDT", Precision: 6, RateToBase: decimal.NewFromFloat(0.999)},
	{Symbol: "WBTC", Precision: 8, RateToBase: decimal.NewFromFloat(57037.22)},
}

type CurrencySvc struct {
	currencyCache map[string]currency
}

func NewCurrencySvc() *CurrencySvc {
	currencyMap := make(map[string]currency)

	for _, c := range currencies {
		currencyMap[c.Symbol] = c
	}

	return &CurrencySvc{currencyCache: currencyMap}
}

func (s CurrencySvc) Rate(from, to string) (decimal.Decimal, error) {
	fromCurrency, ok := s.currencyCache[from]
	if !ok {
		return decimal.Zero, errors.New("invalid 'from' currency")
	}

	toCurrency, ok := s.currencyCache[to]
	if !ok {
		return decimal.Zero, errors.New("invalid 'to' currency")
	}
	rate := fromCurrency.RateToBase.Div(toCurrency.RateToBase).Truncate(maxPrecision)
	return rate, nil
}

func (s CurrencySvc) Exchange(from, to string, amount decimal.Decimal) (decimal.Decimal, error) {
	fromCurrency, ok := s.currencyCache[from]
	if !ok {
		return decimal.Zero, errors.New("invalid 'from' currency")
	}

	toCurrency, ok := s.currencyCache[to]
	if !ok {
		return decimal.Zero, errors.New("invalid 'to' currency")
	}

	amountInBase := amount.Mul(fromCurrency.RateToBase)
	convertedAmount := amountInBase.Div(toCurrency.RateToBase)
	convertedAmount = convertedAmount.Truncate(toCurrency.Precision)

	return convertedAmount, nil
}
