package domain 

type Game interface {
	Save()
	Load(savePoint int) GameSave 
	Play(data GameSave) 
}