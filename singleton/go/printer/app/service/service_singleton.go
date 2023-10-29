package service

import (
	"fmt"
	"printer/domain"
	"sync"
)

var SingletonPrinter *domain.Printer

var lock = &sync.Mutex{}

func NewGetPrinterSingleton_Way3() *domain.Printer {
	if SingletonPrinter == nil {
		lock.Lock()
		defer lock.Unlock()
		if SingletonPrinter == nil {
			SingletonPrinter = NewPrinter(domain.SAMSUNG_PRINTER)
			fmt.Println("Create new printer")
			return SingletonPrinter
		} else {
			fmt.Println("Already printer was created")
			return SingletonPrinter
		}
	} else {
		fmt.Println("Already printer was created")
		return SingletonPrinter
	}
}

var once sync.Once

func NewGetPrintSingleton_Way2() *domain.Printer {
	once.Do(func() {
		SingletonPrinter = NewPrinter(domain.SAMSUNG_PRINTER)
	})

	return SingletonPrinter
}

func NewGetPrinterNotSingleton_Way1() *domain.Printer {
	if SingletonPrinter == nil {
		SingletonPrinter = NewPrinter(domain.SAMSUNG_PRINTER)
		return SingletonPrinter
	} else {
		return SingletonPrinter
	}
}

func NewGetPrinterNotSingleton() *domain.Printer {
	return NewPrinter(domain.SAMSUNG_PRINTER)
}
