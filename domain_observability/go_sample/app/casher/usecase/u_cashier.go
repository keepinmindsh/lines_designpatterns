package usecase

import (
	"github.com/go_sample/domain/cacher"
	"github.com/go_sample/domain/cart"
	"github.com/go_sample/domain/payment"
)

type Service struct {
	RateRule     payment.PriceRate
	ShoppingCart cart.Cart
}

func (s Service) LookUpDiscount(code payment.DiscountCode) payment.RateOptions {
	//TODO implement me
	panic("implement me")
}

func (s Service) ApplyToProducts() float64 {
	//TODO implement me
	panic("implement me")
}

func NewCashier(shoppingCart cart.Cart, priceRate payment.PriceRate) cacher.Cashier {
	return &Service{
		ShoppingCart: shoppingCart,
		RateRule:     priceRate,
	}
}
