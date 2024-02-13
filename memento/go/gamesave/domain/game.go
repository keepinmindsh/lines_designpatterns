package domain

type Game interface {
	Play(data GameSave)
	Save(data GameSave)
	Load() GameSave
}
