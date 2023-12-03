package element

import (
	"compiler/domain"
	"fmt"
)

type Node struct {
}

func (n Node) Process() {
	fmt.Println("Loading node is processing")
}

func NewNode() domain.Element {
	return &Node{}
}
