package lines.service.factory;

import lines.model.Unit;
import lines.model.UnitAction;
import lines.service.usecase.Marine;

public class UnitFactory {

    public static Unit CreateUnit(UnitAction action) {
        return new Marine(action);
    }
}
