<<<<<<< HEAD
package domain 

type Game interface {
	Save()
	Load(savePoint int) GameSave 
	Play(data GameSave) 
}
=======
package domain

type GameData struct {
	Position PlayerPosition
	Status   CharacterStatus
	Story    StoryPoint
}

type PlayerPosition struct {
	X int
	Y int
}

type CharacterStatus struct {
	Health   int
	Level    int
	Strength int
}

type StoryPoint struct {
	StoryAddress string
}

type Game interface {
	Play()
	Save(data GameData)
	Load() GameData
}
>>>>>>> 2eef498 (added for meidater and memento)
