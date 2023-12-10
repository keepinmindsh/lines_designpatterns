package domain

type (
	Cell interface {
		Draw()
	}

	Text interface {
		Draw()
	}
)

type TextType string

const (
	A TextType = "A"
	B TextType = "B"
	C TextType = "C"
	D TextType = "D"
	E TextType = "E"
)

func (t TextType) String() string {
	return string(t)
}
