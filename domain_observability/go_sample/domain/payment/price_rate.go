package payment

type RateOptions struct {
	RatePercent float64
}

type PriceRate interface {
	GetRateRule() RateOptions
}
