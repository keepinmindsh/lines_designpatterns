package lines.service.usecase;

import lines.model.Unit;
import lines.model.UnitAction;

public class Marine implements Unit {

    private final UnitAction unitAction;
    private int accumulatedMineral;

    public Marine(UnitAction unitAction){
        this.unitAction = unitAction;
    }

    @Override
    public int GerMineralCapacity() {
        return accumulatedMineral;
    }

    @Override
    public void SetMineralCapacity(int accumulatedMineral) {
        this.accumulatedMineral = accumulatedMineral;
    }

    @Override
    public void Harvest() {
        this.unitAction.Harvest();
    }

    @Override
    public void Attack() {
        this.unitAction.Attack();
    }

    @Override
    public void Building() {
        this.unitAction.Building();
    }
}
