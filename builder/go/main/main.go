package main

import (
	"builder/domain"
	"fmt"
)

func main() {
	builder := domain.NewReservationBuilder()

	builder.Type(domain.Expedia)

	reservationBuilder := builder.Make()

	fmt.Println("Reservation Type: " + reservationBuilder.RsvnType)
}
