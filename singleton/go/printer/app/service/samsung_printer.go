package service

import "printer/domain"

type samsungPrinter struct {
}

func (s samsungPrinter) PutPaperIn() {
	//TODO implement me
	panic("implement me")
}

func (s samsungPrinter) Print() {
	//TODO implement me
	panic("implement me")
}

func NewSamsungPrinter() domain.Printer {
	return &samsungPrinter{}
}
