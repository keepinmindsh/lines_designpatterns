package parser

import "webserver/domain"

type jsonParser struct {
}

func (j jsonParser) Parse() (interface{}, bool) {
	return nil, false
}

func NewJsonParser() domain.Parser {
	return jsonParser{}
}
