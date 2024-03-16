package main

import (
	"printer/app/action"
	"printer/app/client"
	"printer/app/input"
	"printer/domain"
)

func main() {
	list := make([]domain.Input, 4)

	list = append(list, input.NewLG())
	list = append(list, input.NewSamsung())
	list = append(list, input.NewEpson())
	list = append(list, input.NewPDF())

	printer := client.NewPrinter(list)

	printer.Add(action.NewCheckPrint())
	printer.Add(action.NewInputPrint())
	printer.Add(action.NewNextPrint())

	printer.Print()
}
