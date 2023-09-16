package app.service.expedia

import domain.Reservation

class Expedia : Reservation {
    override fun SchedulePeriod() {
        print("Expedia Scheduled!")
    }

    override fun Book() {
        print("Expedia Booked!")
    }

    override fun PayMoney() {
        print("Expedia Pay Money!")
    }
}