package usecase

import "go_sample/domain"

type ShoppingCart struct {
}

func NewShoppingCart() domain.ShoppingCart {
	return &ShoppingCart{}
}

func (s ShoppingCart) ApplyDiscountCode(code domain.DiscountCode) {

}
