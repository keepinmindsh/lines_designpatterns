package main

import (
	"human_walk/app/human"
	"human_walk/domain"
)

func main() {
	aPerson := human.NewHuman()

	aPerson.HealthStatus(domain.HEALTHY)
	aPerson.Walk()
	aPerson.HealthStatus(domain.SICK)
	aPerson.Walk()
	aPerson.HealthStatus(domain.BROKEN_ANKLE)
	aPerson.Walk()
	aPerson.HealthStatus(domain.BROKEN_BONES)
	aPerson.Walk()
}
