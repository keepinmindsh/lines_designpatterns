package service

import "printer/domain"

func NewPrinter(printType domain.PrintType) domain.Printer {
	switch printType {
	case domain.SAMSUNG_PRINTER :
		return NewSamsungPrinter()
	case domain.APPLE_PRINTER
		return
	default:
		return nil
	}
}

