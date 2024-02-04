package webserver

import "webserver/domain"

type WebServer struct {
	domain.HttpExecutor
}

func NewWebServer(executor domain.HttpExecutor) domain.Processor {
	return WebServer{
		executor,
	}
}
