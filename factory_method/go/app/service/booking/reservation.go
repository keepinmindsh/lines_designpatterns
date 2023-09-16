package booking

import (
	"factory_method/domain"
	"fmt"
)

type Reservation struct{}

func NewBookingReservation() domain.Reservation {
	return &Reservation{}
}

func (r Reservation) SchedulePeriod() {
	fmt.Println("Booking.com scheduled!")
}

func (r Reservation) Book() {
	fmt.Println("Booking.com booked!")
}

func (r Reservation) PayMoney() {
	fmt.Println("Booking.com paymoney!")
}
