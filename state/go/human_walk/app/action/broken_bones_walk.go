package action

import "human_walk/domain"

type brokenBonesWalk struct {
}

func (b brokenBonesWalk) Walk() {
	//TODO implement me
	panic("implement me")
}

func NewBrokenBonesWalk() domain.Walk {
	return brokenBonesWalk{}
}
