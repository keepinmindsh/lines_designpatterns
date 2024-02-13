<<<<<<< HEAD
package domain 


type GameSave struct {
	Position CharacterPosition
	Status CharaterStatus
	Address StoryPoint
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
=======
package domain

type Save interface {
	Save()
	Load()
}
>>>>>>> 2eef498 (added for meidater and memento)
