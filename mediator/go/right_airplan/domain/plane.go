package domain

type AirPlane interface {
	GetName() string
	Landing()
	LandingEnable() bool
	SetLandingEnable(bool)
	Stop()
	PlaneLanded() bool
}
