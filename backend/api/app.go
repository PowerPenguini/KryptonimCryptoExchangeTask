package api

type App struct {
	CurrencySvc CurrencySvc
}

func NewApp(cs CurrencySvc) *App {
	return &App{cs}
}
