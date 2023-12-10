package usecase

import (
	"excel_sheet/domain"
	"excel_sheet/domain/position"
	"fmt"
)

type Cell struct {
	Position *position.Position
	Text     *domain.Text
}

func NewCell(position *position.Position, text *domain.Text) domain.Cell {
	c := &Cell{
		Position: position,
		Text:     text,
	}

	return c
}

func (c *Cell) Draw() {
	if c.Text != nil {

		print(fmt.Sprintf(" Text: Address - %v   ", &(*c.Text)))
		print(fmt.Sprintf(" Text: Pointer Address - %v   ", &c.Text))

		(*c.Text).Draw()
	}
}
