package airstrip

import "airplane/domain"

type airstrip struct {
	controlTower domain.ControlTower
}

func (a airstrip) Landing() {
	a.controlTower.SetStatusAirStrip(true)
}

func NewAirStrip() domain.AirStrip {
	return &airstrip{}
}
