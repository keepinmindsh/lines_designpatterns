package domain

type (
	Chat interface {
		RunContext(message Message)
		MessageCreate()
	}

	Message interface {
		Create()
		Notify()
		Send()
	}
)
