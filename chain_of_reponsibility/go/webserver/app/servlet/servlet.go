package servlet

import "webserver/domain"

type Servlet struct {
	processor domain.Processor
	executors []domain.Executor
}

func MustNewServlet(executor domain.Processor, executors []domain.Executor) Servlet {
	return Servlet{processor: executor, executors: executors}
}

func (s *Servlet) Process(request domain.Request) {
	proxyLength := len(s.executors)

	for i := 0; i < proxyLength; i++ {
		s.executors[i].Pre(request)
	}

	response := s.processor.Do(request)

	for i := proxyLength; i > 0; i-- {
		data, ok := response.(domain.Response)
		if ok {
			s.executors[i-1].Post(data)
		}
	}
}
