package usecase

import (
	"github.com/go_sample/domain/cart"
)

type ShoppingCart struct {
}

func (s ShoppingCart) GetProducts() {
	//TODO implement me
	panic("implement me")
}

func NewShoppingCart() cart.Cart {
	return &ShoppingCart{}
}
