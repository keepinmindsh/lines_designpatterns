package play 

import (
	"gamesave/domain"
)

type GamePlay struct {
	GameData domain.GameSave,
	Save domain.Save
}

func (g *GamePlay) Save() {
	g.Save.Save(g.GameData)
}

func (g *GamePlay) Load(savePoint int) domain.GameSave {
	return g.Save.Load(savePoint)
}

func (g *GamePlay) Play(data domain.GameSave) {
	g.GameData = data
}

func NewGamePlay(save domian.GameSave) domain.Game {
	return &GamePlay{
		Save : save
	}
}