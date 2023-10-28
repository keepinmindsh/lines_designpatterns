package domain

type PrintType string

const (
	SAMSUNG_PRINTER PrintType = "Samsung Printer"
	APPLE_PRINTER   PrintType = "Apple Printer"
)

type Printer interface {
	PutPaperIn()
	Print()
}
