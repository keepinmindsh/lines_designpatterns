package lines.model;

public interface Unit {
    int GerMineralCapacity();
    void SetMineralCapacity(int accumulatedMineral);

    void Harvest();
    void Attack();
    void Building();
}
