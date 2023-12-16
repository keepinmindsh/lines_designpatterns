package cacher

import "github.com/go_sample/domain/payment"

type Cashier interface {
	LookUpDiscount(code payment.DiscountCode) payment.RateOptions
	ApplyToProducts() float64
}
