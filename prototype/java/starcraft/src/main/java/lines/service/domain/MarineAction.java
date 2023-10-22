package lines.service.domain;

import lines.model.UnitAction;

public class MarineAction implements UnitAction {
    @Override
    public void Harvest() {
        System.out.println("수확합니다.");
    }

    @Override
    public void Attack() {
        System.out.println("공격합니다.");
    }

    @Override
    public void Building() {
        System.out.println("건설합니다.");
    }
}
