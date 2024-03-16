package action

import "human_walk/domain"

type brokenAnkleWalk struct {
}

func (b brokenAnkleWalk) Walk() {
	//TODO implement me
	panic("implement me")
}

func NewBrokenAnkleWalk() domain.Walk {
	return brokenAnkleWalk{}
}
