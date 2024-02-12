package domain

type Strip interface {
	SetAirPlanName(string)
	GetAirPlanName() string
	LandingProcedure()
	Complete()
	GetRoadStatus() string
}
