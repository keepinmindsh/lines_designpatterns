package input

import "printer/domain"

type lg struct {
}

func (l *lg) Accept(printer domain.Printer) {
	printer.ExecuteForLG()
}

func NewLG() domain.Input {
	return &lg{}
}
