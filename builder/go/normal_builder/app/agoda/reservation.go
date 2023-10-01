package agoda

import "builder/domain"

type Agoda struct {
}

func NewAgodaReservation() domain.Reservation {
	return &Agoda{}
}

func (a Agoda) MakeReservation(reservation *domain.ReservationBuilder) {
	//TODO implement me
	panic("implement me")
}
