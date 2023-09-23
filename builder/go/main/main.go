package main

import (
	"builder/app/factory"
	"builder/domain"
	"fmt"
)

func main() {
	builder := domain.NewReservationBuilder()

	builder.Type(domain.Expedia).RsvnName("Bongs")

	reservationBuilder := builder.Make()

	reservationWay := factory.NewReservationWay(domain.Expedia)

	reservationWay.MakeReservation(reservationBuilder)

	fmt.Println("Reservation Type: " + reservationBuilder.RsvnType)
}
