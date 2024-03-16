package input

import "printer/domain"

type epson struct {
}

func (e *epson) Accept(printer domain.Printer) {
	printer.ExecuteForEpson()
}

func NewEpson() domain.Input {
	return &epson{}
}
