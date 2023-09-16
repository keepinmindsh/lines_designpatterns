package usecases

import (
	"factory_method/app/service/agoda"
	"factory_method/app/service/booking"
	"factory_method/app/service/expedia"
	"factory_method/domain"
)

func NewReservation(rsvnType domain.ReservationType) domain.Reservation {
	switch rsvnType {
	case domain.Agoda:
		return agoda.NewAgodaReservation()
	case domain.Expedia:
		return expedia.NewExpediaReservation()
	case domain.Booking:
		return booking.NewBookingReservation()
	}

	return nil
}
