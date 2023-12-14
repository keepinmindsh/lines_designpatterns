package service

import "go_sample/domain"

type Service struct {
}

func (s Service) LookUpDiscount(code domain.DiscountCode) {

}

func (s Service) ApplyToCard() float64 {
	return 0.0
}

func NewService() domain.Shopping {
	return &Service{}
}
