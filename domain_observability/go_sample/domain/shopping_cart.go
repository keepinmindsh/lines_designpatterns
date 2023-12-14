package domain

type DiscountCode string

const (
	BlackFridayCoupon DiscountCode = "AZ001"
)

type ShoppingCart interface {
	ApplyDiscountCode(code DiscountCode)
}

type Shopping interface {
	LookUpDiscount(code DiscountCode)
	ApplyToCard() float64
}
