package action

import "human_walk/domain"

type sickWalk struct {
}

func (s sickWalk) Walk() {
	//TODO implement me
	panic("implement me")
}

func NewSickWalk() domain.Walk {
	return sickWalk{}
}
