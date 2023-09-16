package expedia

import (
	"factory_method/domain"
	"fmt"
)

type Reservation struct{}

func NewExpediaReservation() domain.Reservation {
	return &Reservation{}
}

func (r Reservation) SchedulePeriod() {
	fmt.Println("Expedia scheduled!")
}

func (r Reservation) Book() {
	fmt.Println("Expedia booked!")
}

func (r Reservation) PayMoney() {
	fmt.Println("Expedia paymoney!")
}
