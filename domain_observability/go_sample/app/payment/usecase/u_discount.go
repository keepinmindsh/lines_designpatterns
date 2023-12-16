package usecase

import (
	"github.com/go_sample/domain/payment"
	"math/big"
)

type Discount struct {
}

func (d Discount) ApplyDiscountRate(code payment.DiscountCode) big.Float {
	//TODO implement me
	panic("implement me")
}

func NewDiscount() payment.Discount {
	return &Discount{}
}
