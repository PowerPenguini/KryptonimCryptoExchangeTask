package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func newGetRates(app *App) gin.HandlerFunc {
	return func(c *gin.Context) {
		currencyParam := c.Query("currencies")
		if currencyParam == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		currencies := strings.Split(currencyParam, ",")
		if len(currencies) < 2 {
			c.Status(http.StatusBadRequest)
			return
		}

		var response []RateResponse
		for i, from := range currencies {
			for j, to := range currencies {
				if i != j {
					rate, err := app.CurrencySvc.Rate(from, to)
					if err != nil {
						c.Status(http.StatusBadRequest)
						return
					}
					response = append(response, RateResponse{
						From: from,
						To:   to,
						Rate: rate,
					})
				}
			}
		}

		c.JSON(http.StatusOK, response)

	}
}
func newGetExchange(app *App) gin.HandlerFunc {
	return func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")
		amountStr := c.Query("amount")

		if from == "" || to == "" || amountStr == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		amount, err := decimal.NewFromString(amountStr)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		convertedAmount, err := app.CurrencySvc.Exchange(from, to, amount)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"from":   from,
			"to":     to,
			"amount": convertedAmount.String(),
		})
	}
}

func NewServer(app *App) *http.Server {
	r := gin.Default()

	r.GET("/rates", newGetRates(app))
	r.GET("/exchange", newGetExchange(app))

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
