package element

import (
	"compiler/domain"
	"fmt"
)

type Parser struct {
}

func (p Parser) Process() {
	fmt.Println("Parser is processing")
}

func NewParser() domain.Element {
	return &Parser{}
}
