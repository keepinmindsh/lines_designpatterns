package factory

import (
	"builder/app/expedia"
	"builder/domain"
)

func NewReservationWay(rsvnType domain.RsvnType) domain.Reservation {

	switch rsvnType {
	case domain.Expedia:
		return expedia.NewExpediaReservation()
	case domain.Booking:
	case domain.Agoda:
	}

	return nil
}
