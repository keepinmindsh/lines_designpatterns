package usecase

import (
	"memory/domain"
	"strconv"
)

type RealObject struct {
}

func (r RealObject) Load() interface{} {
	var stringList []string
	for i := 0; i < 100; i++ {
		stringList = append(stringList, "Value"+strconv.Itoa(i))
	}

	return stringList
}

func NewRealObject() domain.Loader {
	return RealObject{}
}
