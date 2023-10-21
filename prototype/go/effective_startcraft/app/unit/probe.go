package unit

import (
	"effective_startcraft/domain"
)

type Probe struct {
	AccumulatedMineral int
	domain.UnitAction
}

func (p *Probe) GerMineralCapacity() int {
	return p.AccumulatedMineral
}

func (p *Probe) SetMineralCapacity(accumulatedMineral int) {
	p.AccumulatedMineral = accumulatedMineral
}

func NewCloneUnit(action domain.UnitAction) domain.Unit {
	cloneObj := &Probe{
		0,
		action,
	}

	return cloneObj
}
