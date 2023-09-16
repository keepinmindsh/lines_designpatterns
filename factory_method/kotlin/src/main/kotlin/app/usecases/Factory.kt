package app.usecases

import app.service.agoda.Agoda
import app.service.booking.Booking
import app.service.expedia.Expedia
import domain.Reservation
import domain.ReservationType

class Factory {
    companion object {
        fun GetReservationWay(rsvnType: ReservationType): Reservation {
            when (rsvnType) {
                ReservationType.Agoda -> return Agoda()
                ReservationType.Booking -> return Booking()
                ReservationType.Expedia -> return Expedia()
                else -> throw java.lang.RuntimeException("Error!")
            }
        }
    }
}