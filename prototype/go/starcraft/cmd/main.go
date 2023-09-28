package main

import "starcraft/app/unit"

func main() {
	clonedUnit := unit.NewCloneUnit()

	clonedUnit.Building()

	clonedUnit.Attack()

	clonedUnit.Harvest()
}
