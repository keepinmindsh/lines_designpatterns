package main

import (
	"gamesave/app/play"
	"gamesave/app/save"
	"gamesave/domain"
)

const Latest int = 0

func main() {
	gamePlay := play.NewGamePlay(save.NewSave())

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 20,
		},
		Status: domain.CharaterStatus{
			Status: "good",
		},
		Address: domain.StoryPoint{
			Address: "story-1-1",
		},
	})

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 22,
		},
		Status: domain.CharaterStatus{
			Status: "good",
		},
		Address: domain.StoryPoint{
			Address: "story-1-1",
		},
	})

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 25,
		},
		Status: domain.CharaterStatus{
			Status: "bad",
		},
		Address: domain.StoryPoint{
			Address: "story-1-2",
		},
	})

	gamePlay.Save()

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 35,
		},
		Status: domain.CharaterStatus{
			Status: "good",
		},
		Address: domain.StoryPoint{
			Address: "story-1-8",
		},
	})

	gamePlay.Save()

	gamePlay.Play(domain.GameSave{
		Position: domain.CharacterPosition{
			X: 10,
			Y: 35,
		},
		Status: domain.CharaterStatus{
			Status: "dead",
		},
		Address: domain.StoryPoint{
			Address: "story-1-20",
		},
	})

	gamePlay.Play(gamePlay.Load(Latest))
}
