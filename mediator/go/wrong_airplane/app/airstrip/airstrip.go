package airstrip

import (
	"sync"
	"wrong_airplane/domain"
)

type airstrip struct {
	controlTower *domain.ControlTower
}

var mw sync.Mutex

func (a airstrip) Landing() {
	mw.Lock()
	(*a.controlTower).SetStatusAirStrip(true)
	mw.Unlock()
}

func NewAirStrip(tower *domain.ControlTower) domain.AirStrip {
	return &airstrip{
		controlTower: tower,
	}
}
