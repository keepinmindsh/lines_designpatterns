package input

import "printer/domain"

type pdf struct {
}

func (p pdf) Accept(printer domain.Printer) {
	printer.ExecuteForPDF()
}

func NewPDF() domain.Input {
	return &pdf{}
}
