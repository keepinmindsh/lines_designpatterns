package lines.application

import lines.service.domain.MarineAction
import lines.service.factory.UnitFactory

fun main(){
    val marineAction = MarineAction()

    for (i in 1..30) {
        val createUnit = UnitFactory.CreateUnit(marineAction)

        createUnit.Attack()

        createUnit.Building()

        createUnit.Harvest()
    }

}