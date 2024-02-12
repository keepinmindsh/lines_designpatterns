package airport

import (
	"fmt"
	"right_airplan/app/plane"
	"right_airplan/app/strip"
	"right_airplan/domain"
	"sync"
	"testing"
	"time"
)

var lock = &sync.Mutex{}

func TestDo(t *testing.T) {
	airStrip := strip.NewStrip()

	airplanes := []*plane.Plane{
		plane.NewPlane("대한항공", &airStrip),
		plane.NewPlane("아시아나 항공", &airStrip),
		plane.NewPlane("제주 항공", &airStrip),
		plane.NewPlane("티웨이 항공", &airStrip),
		plane.NewPlane("동방항공", &airStrip),
	}

	lock := &sync.Mutex{}

	var wg sync.WaitGroup

	for {
		for i := 0; i < len(airplanes); i++ {
			if !airplanes[i].PlaneLanded() {
				wg.Add(1)

				go func(item domain.AirPlane) {
					if airStrip.GetRoadStatus() == "BEFORE" || airStrip.GetRoadStatus() == "COMPLETE" {
						if !item.PlaneLanded() {
							if airStrip.GetRoadStatus() != "RUNNING" {
								lock.Lock()
								defer lock.Unlock()
								if airStrip.GetRoadStatus() != "RUNNING" {
									airStrip.LandingProcedure()

									item.Landing()

									item.Stop()
								}
							}
						}
					}

					wg.Done()
				}(airplanes[i])
			}
		}

		var waitingAirplane string
		for _, airplane := range airplanes {
			if !airplane.LandingEnable() && !airplane.PlaneLanded() {
				waitingAirplane += "[" + airplane.GetName() + "]"
			}
		}
		fmt.Println(waitingAirplane + " 착륙 대기중")

		var isFinished bool
		for _, airplane := range airplanes {
			if !airplane.PlaneLanded() {
				isFinished = false
				break
			}

			isFinished = true
		}

		if isFinished {
			break
		}
		time.Sleep(time.Second * 1)
	}

	wg.Wait()
}
