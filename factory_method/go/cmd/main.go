package main

import (
	"factory_method/app/usecases"
	"factory_method/domain"
)

func main() {
	reservation := usecases.NewReservation(domain.Agoda)

	reservation.SchedulePeriod()

	reservation.Book()

	reservation.PayMoney()
}
