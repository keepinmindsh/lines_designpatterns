package parser

import "webserver/domain"

type stringParser struct {
}

func (s stringParser) Parse() (interface{}, bool) {
	return nil, true
}

func NewStringParser() domain.Parser {
	return stringParser{}
}
