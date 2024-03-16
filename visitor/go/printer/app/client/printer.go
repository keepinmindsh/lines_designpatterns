package client

import (
	"printer/domain"
)

type Printer struct {
	PrintInput []domain.Input
}

var printAction []domain.Printer

func (p *Printer) Add(printerAction domain.Printer) {
	printAction = append(printAction, printerAction)
}

func (p *Printer) Print() {
	for _, input := range p.PrintInput {
		for _, printer := range printAction {
			input.Accept(printer)
		}
	}
}

func NewPrinter(inputList []domain.Input) Printer {
	return Printer{
		PrintInput: inputList,
	}
}
