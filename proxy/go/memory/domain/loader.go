package domain

type LoaderType string

const (
	META LoaderType = "meta"
	REAL LoaderType = "real"
)

type Loader interface {
	Load() interface{}
}
