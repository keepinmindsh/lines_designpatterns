package app.service.agoda

import domain.Reservation

class Agoda : Reservation {
    override fun SchedulePeriod() {
        print("Agoda Scheduled!")
    }

    override fun Book() {
        print("Agoda Booked!")
    }

    override fun PayMoney() {
        print("Agoda Pay Money!")
    }
}