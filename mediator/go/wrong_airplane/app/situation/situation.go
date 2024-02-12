package situation

import (
	"fmt"
	"sync"
	"time"
	"wrong_airplane/app/queue"
	"wrong_airplane/domain"
)

type Situation struct {
	strip     domain.AirStrip
	airPlanes []domain.AirPlane
	tower     *domain.ControlTower
	queue     *queue.Queue
}

func NewLandingSituation(airPlanes []domain.AirPlane) *Situation {
	return &Situation{
		airPlanes: airPlanes,
		queue:     queue.NewQueue(),
	}
}

func (s *Situation) Do() {
	var wg sync.WaitGroup

	wg.Add(len(s.airPlanes))
	for {
		for _, plane := range s.airPlanes {

			if s.queue.GetCount() == 0 {
				s.queue.Add(&plane)
			}

			go func(item domain.AirPlane) {
				if s.queue.CanBeRun() && !item.Landed() {
					deQueue := s.queue.DeQueue()
					(*deQueue).Takeoff()

					s.queue.Remove()

					wg.Done()
				}
			}(plane)
		}
		time.Sleep(time.Second * 2)
		fmt.Println("Airplane is landing...")

		isDone := true
		for _, item := range s.airPlanes {
			if !item.Landed() {
				isDone = false
				break
			}
		}

		if isDone {
			break
		}
	}

	wg.Wait()
}
