package application;

import lines.model.Unit;
import lines.model.UnitAction;
import lines.service.domain.MarineAction;
import lines.service.factory.UnitFactory;

public class Application {
    public static void main(String[] args) {
        UnitAction unitAction = new MarineAction();

        for (int i = 0; i < 60; i++) {
            Unit unit = UnitFactory.CreateUnit(unitAction);

            unit.Attack();

            unit.Building();

            unit.Harvest();
        }
    }
}
