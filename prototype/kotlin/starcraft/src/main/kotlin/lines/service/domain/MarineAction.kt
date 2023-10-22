package lines.service.domain

import lines.model.UnitAction

class MarineAction : UnitAction {
    override fun Harvest() {
        print("수락합니다.")
    }

    override fun Attack() {
        print("공격합니다.")
    }

    override fun Building() {
        print("건설합니다.")
    }
}