package main

import (
	"effective_startcraft/app/action"
	"effective_startcraft/app/building"
	"sync"
)

func main() {
	nexus := building.NewNexus()

	action := action.NewAction()

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {
			clonedUnit := nexus.CreateUnit(action)

			clonedUnit.Building()

			clonedUnit.Attack()

			for i := 0; i < 60; i++ {
				harvest := clonedUnit.Harvest(clonedUnit.GerMineralCapacity())

				clonedUnit.SetMineralCapacity(harvest)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
