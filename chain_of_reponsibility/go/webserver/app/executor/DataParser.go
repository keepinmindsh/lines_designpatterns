package executor

import "webserver/domain"

type DataParser struct {
	parserChain []domain.Parser
}

func (d DataParser) Do(request domain.Request) interface{} {
	for _, parser := range d.parserChain {
		if _, ok := parser.Parse(); ok {
			return domain.Response{Data: "Hi"}
		}
	}

	return nil
}

func NewDataParser(parserChain []domain.Parser) domain.HttpExecutor {
	return DataParser{
		parserChain: parserChain,
	}
}
