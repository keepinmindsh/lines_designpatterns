package service

import "printer/domain"

type applePrinter struct {
}

func (a applePrinter) PutPaperIn() {
	//TODO implement me
	panic("implement me")
}

func (a applePrinter) Print() {
	//TODO implement me
	panic("implement me")
}

func NewApplePrinter() domain.Printer {
	return &applePrinter{}
}
