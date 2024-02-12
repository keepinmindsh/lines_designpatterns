package queue

import (
	"sync"
	"wrong_airplane/domain"
)

type Queue struct {
	blockingQueue []*domain.AirPlane
}

var mw sync.Mutex

func NewQueue() *Queue {
	return &Queue{
		blockingQueue: []*domain.AirPlane{},
	}
}

func (q *Queue) Add(plane *domain.AirPlane) {
	mw.Lock()
	if len(q.blockingQueue) == 0 {
		q.blockingQueue = append(q.blockingQueue, plane)
	}
	mw.Unlock()
}

func (q *Queue) GetCount() int {
	return len(q.blockingQueue)
}

func (q *Queue) DeQueue() *domain.AirPlane {
	mw.Lock()
	ret := make([]*domain.AirPlane, 0)
	item := q.blockingQueue[0]
	ret = append(ret, q.blockingQueue[:0]...)
	q.blockingQueue = append(ret, q.blockingQueue[0+1:]...)
	mw.Unlock()

	return item
}

func (q *Queue) CanBeRun() bool {
	return len(q.blockingQueue) > 0
}

func (q *Queue) Remove() {
	mw.Lock()
	ret := make([]*domain.AirPlane, 0)
	q.blockingQueue = append(ret, q.blockingQueue[:0]...)
	mw.Unlock()
}
