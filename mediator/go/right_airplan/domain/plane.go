package domain

type AirPlane interface {
	GetName() string
	Landing()
	LandingEnable() bool
<<<<<<< HEAD
	SetLandingEnable(bool)
=======
>>>>>>> 2eef498 (added for meidater and memento)
	Stop()
	PlaneLanded() bool
}
