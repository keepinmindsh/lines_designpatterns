package strip

import (
	"right_airplan/domain"
)

type strip struct {
	AirPlanName string
}

func (s *strip) SetAirPlanName(airPlaneName string) {
	s.AirPlanName = airPlaneName
}

func (s *strip) GetAirPlanName() string {
	return s.AirPlanName
}

func (s *strip) LandingProcedure() {
	s.AirPlanName = "RUNNING"
}

func (s *strip) Complete() {
	s.AirPlanName = "COMPLETE"
}

func (s *strip) GetRoadStatus() string {
	return s.AirPlanName
}

func NewStrip() domain.Strip {
	return &strip{
		AirPlanName: "BEFORE",
	}
}
