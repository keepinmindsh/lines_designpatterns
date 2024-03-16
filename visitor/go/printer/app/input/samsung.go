package input

import "printer/domain"

type samsung struct {
}

func (s *samsung) Accept(printer domain.Printer) {
	printer.ExecuteForSamsung()
}

func NewSamsung() domain.Input {
	return &samsung{}
}
