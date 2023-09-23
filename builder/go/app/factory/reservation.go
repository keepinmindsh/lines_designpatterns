package factory

import "builder/domain"

func NewReservationWay(rsvnType domain.RsvnType) domain.Reservation {

	switch rsvnType {
	case domain.Expedia:
	case domain.Booking:
	case domain.Agoda:
	}

	return nil
}
