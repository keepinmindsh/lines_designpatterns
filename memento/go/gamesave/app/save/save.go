package save

import (
	"gamesave/domain"
)

type Save struct {
	SaveData []domain.GameSave
}

func (s *Save) Save(data domain.GameSave) {
	s.SaveData = append(s.SaveData, data)
}

func (s *Save) Load(savePoint int) domain.GameSave {
	if s.SaveData[savePoint] != (domain.GameSave{}) {
		return s.SaveData[savePoint]
	} else {
		return domain.GameSave{}
	}
}

func NewSave() domain.Save {
	return &Save{
		SaveData: []domain.GameSave{},
	}
}
