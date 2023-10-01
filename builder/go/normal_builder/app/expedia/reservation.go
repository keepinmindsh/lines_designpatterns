package expedia

import "builder/domain"

type Expedia struct {
}

func NewExpediaReservation() domain.Reservation {
	return &Expedia{}
}

func (e Expedia) MakeReservation(reservation *domain.ReservationBuilder) {
	//TODO implement me
	panic("implement me")
}
