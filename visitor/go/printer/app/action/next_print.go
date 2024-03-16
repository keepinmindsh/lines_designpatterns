package action

import "printer/domain"

type nextPrint struct{}

func (n nextPrint) ExecuteForEpson() {
	//TODO implement me
	panic("implement me")
}

func (n nextPrint) ExecuteForPDF() {
	//TODO implement me
	panic("implement me")
}

func (n nextPrint) ExecuteForSamsung() {
	//TODO implement me
	panic("implement me")
}

func (n nextPrint) ExecuteForLG() {
	//TODO implement me
	panic("implement me")
}

func NewNextPrint() domain.Printer {
	return &nextPrint{}
}
