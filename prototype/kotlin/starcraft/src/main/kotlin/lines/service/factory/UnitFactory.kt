package lines.service.factory

import lines.model.UnitAction
import lines.model.Unit
import lines.service.usecase.Marine

class UnitFactory {
    companion object {
        fun CreateUnit(unitAction: UnitAction) : Unit {
            return Marine(unitAction, 50)
        }
    }
}