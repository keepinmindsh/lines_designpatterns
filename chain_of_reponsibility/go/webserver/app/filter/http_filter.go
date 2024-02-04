package filter

import (
	"fmt"
	"webserver/domain"
)

type httpFilter struct {
}

func (h httpFilter) Pre(request domain.Request) {
	fmt.Println("Pre Http Filter")
}

func (h httpFilter) Post(response domain.Response) {
	fmt.Println("Post Http Filter")
}

func NewFilter() domain.Executor {
	return httpFilter{}
}
