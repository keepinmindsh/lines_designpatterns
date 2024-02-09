package airplane

import (
	"airplane/domain"
	"fmt"
	"time"
)

type airplane struct {
	Name         string
	IsLanded     bool
	controlTower domain.ControlTower
	strip        domain.AirStrip
}

var TakeOffTime = 30

func (a *airplane) Takeoff() {
	for i := 0; i < TakeOffTime; i++ {
		time.Sleep(time.Second * 1)
	}
	a.strip.Landing()

	a.IsLanded = true
}

func (a *airplane) Moving() {
	fmt.Println("Moving " + a.Name)
}

func (a *airplane) CheckPossibleLanding() bool {
	isPass := false

	if a.controlTower.IsLandingEnable() {
		isPass = true
	}

	isPass = false
	if a.controlTower.IsAirStripEmpty() {
		isPass = true
	}

	return isPass
}

func NewAirPlane(airPlane string, isLanded bool) domain.AirPlane {
	return &airplane{
		Name:     airPlane,
		IsLanded: isLanded,
	}
}
