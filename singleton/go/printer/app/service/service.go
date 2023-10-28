package service

import (
	"printer/domain"
	"sync"
)

var Printer domain.Printer

func NewGetPrinter() domain.Printer {
	once := sync.Once{}

	if Printer != nil {
		once.Do(func() {
			Printer = NewPrinter(domain.SAMSUNG_PRINTER)
		})
	} else {
		return Printer
	}

	return nil
}
