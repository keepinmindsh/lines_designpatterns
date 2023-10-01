package booking

import "builder/domain"

type Booking struct {
}

func NewBookingReservation() domain.Reservation {
	return &Booking{}
}

func (b Booking) MakeReservation(reservation *domain.ReservationBuilder) {
	//TODO implement me
	panic("implement me")
}
