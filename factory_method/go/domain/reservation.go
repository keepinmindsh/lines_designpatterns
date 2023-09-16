package domain

type ReservationType string

const (
	Agoda   ReservationType = "Agoda"
	Booking ReservationType = "Booking"
	Expedia ReservationType = "Expedia"
)

type (
	Reservation interface {
		SchedulePeriod()
		Book()
		PayMoney()
	}
)
