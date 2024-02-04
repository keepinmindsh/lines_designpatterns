package domain

type Processor interface {
	HttpExecutor
}

type Executor interface {
	Pre(Request)
	Post(Response)
}

type HttpExecutor interface {
	Do(request Request) interface{}
}
