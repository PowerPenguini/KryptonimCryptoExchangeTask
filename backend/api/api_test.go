package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"kryptonim/api"
	"kryptonim/infra"
)

func TestGetRates(t *testing.T) {
	gin.SetMode(gin.TestMode)
	currencySvc := infra.NewCurrencySvc()
	app := api.NewApp(currencySvc)
	server := api.NewServer(app)

	req, _ := http.NewRequest("GET", "/rates?currencies=GATE,USD,BEER", nil)
	w := httptest.NewRecorder()
	server.Handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `[{"from":"GATE","to":"USD","rate":"6.87"},{"from":"GATE","to":"BEER","rate":"279154.81511580658269"},{"from":"USD","to":"GATE","rate":"0.1455604075691412"},{"from":"USD","to":"BEER","rate":"40633.8886631450629825"},{"from":"BEER","to":"GATE","rate":"0.0000035822416303"},{"from":"BEER","to":"USD","rate":"0.00002461"}]`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestGetRatesBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	currencySvc := infra.NewCurrencySvc()
	app := api.NewApp(currencySvc)
	server := api.NewServer(app)

	req, _ := http.NewRequest("GET", "/rates?currencies=USD", nil)
	w := httptest.NewRecorder()
	server.Handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	assert.Empty(t, w.Body.String())
}
