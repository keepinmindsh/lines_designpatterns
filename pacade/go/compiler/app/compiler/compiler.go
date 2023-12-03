package compiler

import (
	"compiler/app/builder"
	"compiler/domain"
)

type Compiler struct {
	element *builder.CompilerElement
}

func (c Compiler) Compile() {
	c.element.GetParser().Process()

	c.element.GetScanner().Process()

	c.element.GetProgramNode().Process()
}

func NewCompiler(element *builder.CompilerElement) domain.Compiler {
	return &Compiler{
		element: element,
	}
}
