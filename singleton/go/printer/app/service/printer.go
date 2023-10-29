package service

import "printer/domain"

func NewPrinter(printType domain.PrintType) *domain.Printer {
	switch printType {
	case domain.SAMSUNG_PRINTER:
		newSamsungPrinter := NewSamsungPrinter()
		return &newSamsungPrinter
	case domain.APPLE_PRINTER:
		newApplePrinter := NewApplePrinter()
		return &newApplePrinter
	default:
		return nil
	}
}
