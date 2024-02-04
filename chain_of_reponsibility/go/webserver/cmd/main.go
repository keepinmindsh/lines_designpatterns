package main

import (
	"fmt"
	"webserver/app/executor"
	"webserver/app/filter"
	"webserver/app/interceptor"
	"webserver/app/parser"
	"webserver/app/servlet"
	"webserver/app/webserver"
	"webserver/domain"
)

func main() {
	serverChain := webserver.NewWebServer(executor.NewDataParser(
		[]domain.Parser{
			parser.NewJsonParser(),
			parser.NewStringParser(),
		}))

	newServlet := servlet.MustNewServlet(serverChain, []domain.Executor{
		filter.NewFilter(),
		interceptor.NewHttpInterceptor(),
	})

	newServlet.Process(domain.Request{
		URL:         "https://lines/hello.do",
		ContentType: "Json",
		Data:        "Im a Json",
	})

	fmt.Println("Finished 1 Thread")

	newServlet.Process(domain.Request{
		URL:         "https://lines/hello",
		ContentType: "MultiPart",
		Data:        "Im a MultiPart",
	})

	fmt.Println("Finished 2 Thread")
}
