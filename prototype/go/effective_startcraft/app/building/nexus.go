package building

import (
	"effective_startcraft/app/unit"
	"effective_startcraft/domain"
)

type Nexus struct{}

func (n Nexus) CreateUnit() domain.Unit {
	return unit.NewCloneUnit()
}

func NewNexus() domain.Building {
	return &Nexus{}
}
