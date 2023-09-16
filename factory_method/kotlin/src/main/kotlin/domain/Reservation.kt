package domain

enum class ReservationType {
    Agoda, Booking, Expedia
}

interface Reservation {
    fun SchedulePeriod()
    fun Book()
    fun PayMoney()
}