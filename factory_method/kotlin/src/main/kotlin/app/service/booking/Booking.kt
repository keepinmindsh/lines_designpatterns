package app.service.booking

import domain.Reservation

class Booking : Reservation{
    override fun SchedulePeriod() {
        print("Booking.com Scheduled!")
    }

    override fun Book() {
        print("Booking.com Booked!")
    }

    override fun PayMoney() {
        print("Booking.com Pay Money")
    }
}