package domain

type AirPlane interface {
	Takeoff()
	Moving()
	CheckPossibleLanding() bool
	Landed() bool
}
