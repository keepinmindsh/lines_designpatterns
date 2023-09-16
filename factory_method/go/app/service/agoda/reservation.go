package agoda

import (
	"factory_method/domain"
	"fmt"
)

type Reservation struct{}

func NewAgodaReservation() domain.Reservation {
	return &Reservation{}
}

func (r Reservation) SchedulePeriod() {
	fmt.Println("Agoda scheduled!")
}

func (r Reservation) Book() {
	fmt.Println("Agoda booked!")
}

func (r Reservation) PayMoney() {
	fmt.Println("Agoda payMoney!")
}
