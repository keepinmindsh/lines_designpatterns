package domain

type GameSave struct {
	Position CharacterPosition
	Status   CharaterStatus
	Address  StoryPoint
}

type CharacterPosition struct {
	X int
	Y int
}

type CharaterStatus struct {
	Status string
}

type StoryPoint struct {
	Address string
}

type Save interface {
	Save(data GameSave)
	Load(savePoint int) GameSave
}
