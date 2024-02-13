package play

import (
	"gamesave/domain"
)

type GamePlay struct {
	GameData   domain.GameSave
	SaveModule domain.Save
}

func (g *GamePlay) Save() {
	g.SaveModule.Save(g.GameData)
}

func (g *GamePlay) Load(savePoint int) domain.GameSave {
	return g.SaveModule.Load(savePoint)
}

func (g *GamePlay) Play(data domain.GameSave) {
	g.GameData = data
}

func NewGamePlay(save domain.Save) domain.Game {
	return &GamePlay{
		SaveModule: save,
	}
}
