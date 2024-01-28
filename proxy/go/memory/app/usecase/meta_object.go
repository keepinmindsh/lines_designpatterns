package usecase

import "memory/domain"

type MetaObject struct {
}

func (m MetaObject) Load() interface{} {
	return "String Array"
}

func NewMetaObject() domain.Loader {
	return MetaObject{}
}
