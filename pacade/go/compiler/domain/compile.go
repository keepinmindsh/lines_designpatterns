package domain

type Compiler interface {
	Compile()
}

type Element interface {
	Process()
}
