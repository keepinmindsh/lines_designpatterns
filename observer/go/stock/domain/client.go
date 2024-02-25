package domain

type MarketData struct {
	Name    string
	Pricing int
}

type Client interface {
	Update(marketData MarketData)
}
