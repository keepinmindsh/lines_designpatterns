package domain

type (
	HealthType string

	Action interface {
		Walk()
		HealthStatus(healthType HealthType)
	}

	Walk interface {
		Walk()
	}
)

const (
	SICK         HealthType = "Sick"
	HEALTHY      HealthType = "Healthy"
	BROKEN_ANKLE HealthType = "BrokenAnkle"
	BROKEN_BONES HealthType = "BrokenBones"
)
