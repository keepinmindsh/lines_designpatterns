package domain

type Stock interface {
	Register(client Client)
	RunStockMarket()
}
