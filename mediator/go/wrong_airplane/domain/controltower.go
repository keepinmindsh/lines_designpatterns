package domain

type ControlTower interface {
	SetStatusLanding(bool)
	SetStatusAirStrip(bool)
	IsLandingEnable() bool
	IsAirStripEmpty() bool
}
