package action

import "printer/domain"

type checkPrint struct{}

func (c checkPrint) ExecuteForEpson() {
	//TODO implement me
	panic("implement me")
}

func (c checkPrint) ExecuteForPDF() {
	//TODO implement me
	panic("implement me")
}

func (c checkPrint) ExecuteForSamsung() {
	//TODO implement me
	panic("implement me")
}

func (c checkPrint) ExecuteForLG() {
	//TODO implement me
	panic("implement me")
}

func NewCheckPrint() domain.Printer {
	return checkPrint{}
}
