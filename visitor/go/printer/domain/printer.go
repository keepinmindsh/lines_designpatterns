package domain

type (
	Printer interface {
		ExecuteForEpson()
		ExecuteForPDF()
		ExecuteForSamsung()
		ExecuteForLG()
	}

	Input interface {
		Accept(printer Printer)
	}
)
