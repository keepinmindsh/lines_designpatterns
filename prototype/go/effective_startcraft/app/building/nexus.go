package building

import (
	"effective_startcraft/app/unit"
	"effective_startcraft/domain"
)

type Nexus struct{}

func (n Nexus) CreateUnit(action domain.UnitAction) domain.Unit {
	return unit.NewCloneUnit(action)
}

func NewNexus() domain.Building {
	return &Nexus{}
}
