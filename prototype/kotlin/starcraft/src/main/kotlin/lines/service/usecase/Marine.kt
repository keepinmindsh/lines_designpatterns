package lines.service.usecase

import lines.model.Unit
import lines.model.UnitAction

class Marine constructor(val unitAction: UnitAction, var mineralCapacity: Int) : Unit {
    override fun GetMineralCapacity(): Int {
        return this.mineralCapacity
    }

    override fun SetMineralCapacity(accumulateMineral: Int) {
        this.mineralCapacity = accumulateMineral
    }

    override fun Harvest() {
        this.unitAction.Harvest()
    }

    override fun Attack() {
        this.unitAction.Attack()
    }

    override fun Building() {
        this.unitAction.Building()
    }
}