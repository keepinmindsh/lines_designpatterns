package airplane

import (
	"fmt"
	"strconv"
	"time"
	"wrong_airplane/domain"
)

type airplane struct {
	Name         string
	IsLanded     bool
	controlTower *domain.ControlTower
	strip        *domain.AirStrip
}

var TakeOffTime = 10

func (a *airplane) Takeoff() {
	for i := 0; i < TakeOffTime; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(a.Name + "is Landing..." + strconv.Itoa(i) + "second")
	}
	(*a.strip).Landing()

	a.IsLanded = true

	(*a.controlTower).SetStatusAirStrip(true)
	(*a.controlTower).SetStatusLanding(true)

	fmt.Println(a.Name + "'s Landing has been completed")
}

func (a *airplane) Moving() {
	fmt.Println("Moving " + a.Name)
}

func (a *airplane) CheckPossibleLanding() bool {
	isPass := false
	if (*a.controlTower).IsLandingEnable() {
		isPass = true
	}

	isPass = false
	if (*a.controlTower).IsAirStripEmpty() {
		isPass = true
	}

	return isPass
}

func (a *airplane) Landed() bool {
	return a.IsLanded
}

func NewAirPlane(airPlane string, isLanded bool, controlTower *domain.ControlTower, strip *domain.AirStrip) domain.AirPlane {
	return &airplane{
		Name:         airPlane,
		IsLanded:     isLanded,
		controlTower: controlTower,
		strip:        strip,
	}
}
