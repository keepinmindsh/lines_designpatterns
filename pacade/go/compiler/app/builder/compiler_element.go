package builder

import "compiler/domain"

type CompilerElement struct {
	scanElement        domain.Element
	parserElement      domain.Element
	programNodeElement domain.Element
}

func NewCompilerBuilder() *CompilerElement {
	return &CompilerElement{}
}

func (c *CompilerElement) Scanner(element domain.Element) *CompilerElement {
	c.scanElement = element
	return c
}

func (c *CompilerElement) Parser(element domain.Element) *CompilerElement {
	c.parserElement = element
	return c
}

func (c *CompilerElement) ProgramNode(element domain.Element) *CompilerElement {
	c.programNodeElement = element
	return c
}

func (c *CompilerElement) GetScanner() domain.Element {
	return c.scanElement
}

func (c *CompilerElement) GetParser() domain.Element {
	return c.parserElement
}

func (c *CompilerElement) GetProgramNode() domain.Element {
	return c.programNodeElement
}
