package domain

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
