package domain

type Strip interface {
	SetAirPlanName(string)
<<<<<<< HEAD
	GetAirPlanName() string
=======
>>>>>>> 2eef498 (added for meidater and memento)
	LandingProcedure()
	Complete()
	GetRoadStatus() string
}
