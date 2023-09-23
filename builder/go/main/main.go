package main

import "builder/domain"

func main() {
	builder := domain.NewReservationBuilder()

	builder.Type(domain.Expedia)
}
