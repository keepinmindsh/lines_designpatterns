package main

import (
	"excel_sheet/app/usecase"
	"excel_sheet/domain"
	"excel_sheet/domain/position"
	"fmt"
)

func main() {

	var text domain.Text
	for i := 0; i < 5; i++ {
		text = usecase.GetText(domain.A)

		cell := usecase.NewCell(&position.Position{
			X: i,
			Y: i,
		}, &text)

		print(fmt.Sprintf("Cell Address : %v   ", &cell))

		cell.Draw()
	}
}
