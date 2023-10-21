package domain

type (
	Unit interface {
		GerMineralCapacity() int
		SetMineralCapacity(accumulatedMineral int)
		UnitAction
	}

	UnitAction interface {
		Harvest(accumulatedMineral int) int
		Attack()
		Building()
	}
)
