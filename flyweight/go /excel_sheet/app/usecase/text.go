package usecase

import "excel_sheet/domain"

type a struct{}

func (a a) Draw() {
	print("  A \r\n")
}

type b struct{}

func (b b) Draw() {
	print("  B \n")
}

type c struct{}

func (c c) Draw() {
	print("  C \n")
}

type d struct{}

func (d d) Draw() {
	print("  D \n")
}

type e struct{}

func (e e) Draw() {
	print("  E \n")
}

var TextA domain.Text
var TextB domain.Text
var TextC domain.Text
var TextD domain.Text
var TextE domain.Text

func GetText(txtType domain.TextType) domain.Text {
	switch txtType {
	case domain.A:
		if TextA != nil {
			return TextA
		} else {
			TextA = a{}
			return TextA
		}
	case domain.B:
		if TextB != nil {
			return TextB
		} else {
			TextB = b{}
			return TextB
		}
	case domain.C:
		if TextC != nil {
			return TextC
		} else {
			TextC = c{}
			return TextC
		}
	case domain.D:
		if TextD != nil {
			return TextD
		} else {
			TextD = d{}
			return TextD
		}
	case domain.E:
		if TextE != nil {
			return TextE
		} else {
			TextE = e{}
			return TextE
		}
	default:
		return nil
	}
}
