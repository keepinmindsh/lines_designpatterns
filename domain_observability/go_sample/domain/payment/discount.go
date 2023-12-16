package payment

import (
	"math/big"
)

type DiscountCode string

const (
	BlackFridayCoupon DiscountCode = "AZ001"
)

type Discount interface {
	ApplyDiscountRate(code DiscountCode) big.Float
}
