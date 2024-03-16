package human

import (
	"human_walk/app/action"
	"human_walk/domain"
)

type human struct {
	HealthyStatus domain.HealthType

	SickWalk        domain.Walk
	HealthyWalk     domain.Walk
	BrokenAnkleWalk domain.Walk
	BrokenBonesWalk domain.Walk
}

func (h *human) HealthStatus(healthStatus domain.HealthType) {
	h.HealthyStatus = healthStatus
}

func (h *human) Walk() {
	switch h.HealthyStatus {
	case domain.SICK:
		h.SickWalk.Walk()
	case domain.HEALTHY:
		h.HealthyWalk.Walk()
	case domain.BROKEN_ANKLE:
		h.BrokenAnkleWalk.Walk()
	case domain.BROKEN_BONES:
		h.BrokenBonesWalk.Walk()
	}
}

func NewHuman() domain.Action {
	return &human{
		SickWalk:        action.NewSickWalk(),
		HealthyWalk:     action.NewHealthyWork(),
		BrokenAnkleWalk: action.NewBrokenAnkleWalk(),
		BrokenBonesWalk: action.NewBrokenBonesWalk(),
	}
}
