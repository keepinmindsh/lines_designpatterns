package situation

import "airplane/domain"

type Situation struct {
	strip     domain.AirStrip
	airPlanes []domain.AirPlane
	tower     domain.ControlTower
}

func NewLandingSituation(strip domain.AirStrip, airPlanes []domain.AirPlane, tower domain.ControlTower) *Situation {
	return &Situation{
		strip:     strip,
		airPlanes: airPlanes,
		tower:     tower,
	}
}

func (s Situation) Do() {
	go func() {
		for {
			for _, plane := range s.airPlanes {
				// todo 고루틴 내에 실제 실시간을 값의 변경을 체크하는 루틴을 추가 고민 및 설계 필요
				go func(item domain.AirPlane) {
					if item.CheckPossibleLanding() {
						item.Takeoff()
					}
				}(plane)
			}
		}
	}()
}
