package building

import (
	"starcraft/app/unit"
	"starcraft/domain"
)

type Nexus struct{}

func (n Nexus) CreateUnit() domain.Unit {
	return unit.NewCloneUnit()
}

func NewNexus() domain.Building {
	return &Nexus{}
}
