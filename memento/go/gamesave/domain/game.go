package domain

type Game interface {
	Play(data GameSave)
	Save()
	Load(SavePoint int) GameSave
}
