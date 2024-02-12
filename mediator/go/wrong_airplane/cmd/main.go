package main

import (
	"wrong_airplane/app/airplane"
	"wrong_airplane/app/airstrip"
	"wrong_airplane/app/controltower"
	"wrong_airplane/app/situation"
	"wrong_airplane/domain"
)

func main() {
	tower := controltower.NewControlTower()
	strip := airstrip.NewAirStrip(&tower)

	landingSituation := situation.NewLandingSituation(
		[]domain.AirPlane{
			airplane.NewAirPlane("Plane 1", false, &tower, &strip),
			airplane.NewAirPlane("Plane 2", false, &tower, &strip),
			airplane.NewAirPlane("Plane 3", false, &tower, &strip),
			airplane.NewAirPlane("Plane 4", false, &tower, &strip),
			airplane.NewAirPlane("Plane 5", false, &tower, &strip),
		})

	landingSituation.Do()
}
