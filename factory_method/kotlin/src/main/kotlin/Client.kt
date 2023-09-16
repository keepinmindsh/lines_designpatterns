import app.usecases.Factory
import domain.ReservationType

fun main(){
    val reservation = Factory.GetReservationWay(ReservationType.Agoda)

    reservation.SchedulePeriod()

    reservation.Book()

    reservation.PayMoney()
}
