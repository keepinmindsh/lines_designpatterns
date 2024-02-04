package domain

type Parser interface {
	Parse() (interface{}, bool)
}
