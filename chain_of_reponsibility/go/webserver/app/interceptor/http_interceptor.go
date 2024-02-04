package interceptor

import (
	"fmt"
	"webserver/domain"
)

type httpInterceptor struct {
}

func (h httpInterceptor) Pre(request domain.Request) {
	fmt.Println("Pre Processing")
}

func (h httpInterceptor) Post(request domain.Response) {
	fmt.Println("Post Processing")
}

func NewHttpInterceptor() domain.Executor {
	return httpInterceptor{}
}
