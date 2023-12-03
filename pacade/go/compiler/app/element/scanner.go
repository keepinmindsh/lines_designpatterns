package element

import (
	"compiler/domain"
	"fmt"
)

type Scanner struct {
}

func (s Scanner) Process() {
	fmt.Println("Scan is processing")
}

func NewScanner() domain.Element {
	return &Scanner{}
}
