package action

import "printer/domain"

type inputPrint struct{}

func (i inputPrint) ExecuteForEpson() {
	//TODO implement me
	panic("implement me")
}

func (i inputPrint) ExecuteForPDF() {
	//TODO implement me
	panic("implement me")
}

func (i inputPrint) ExecuteForSamsung() {
	//TODO implement me
	panic("implement me")
}

func (i inputPrint) ExecuteForLG() {
	//TODO implement me
	panic("implement me")
}

func NewInputPrint() domain.Printer {
	return inputPrint{}
}
