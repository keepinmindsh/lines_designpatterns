package action

import "human_walk/domain"

type healthyWork struct {
}

func (h healthyWork) Walk() {
	//TODO implement me
	panic("implement me")
}

func NewHealthyWork() domain.Walk {
	return healthyWork{}
}
