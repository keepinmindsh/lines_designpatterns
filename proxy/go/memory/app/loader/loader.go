package loader

import (
	"memory/app/usecase"
	"memory/domain"
	"sync"
)

type ObjectLoader struct {
	LoaderType domain.LoaderType
	MetaLoader domain.Loader
}

var objectSync sync.Once

var cacheValue []string

func (o ObjectLoader) Load() interface{} {
	if o.LoaderType == domain.META {
		return o.MetaLoader.Load()
	}

	if len(cacheValue) == 0 {
		if o.LoaderType == domain.REAL {
			objectSync.Do(func() {
				object := usecase.NewRealObject().Load().([]string)
				cacheValue = object
			})
		}
	}

	return cacheValue
}

func MustNewObjectLoader(loaderType domain.LoaderType) domain.Loader {
	return ObjectLoader{LoaderType: loaderType, MetaLoader: usecase.NewMetaObject()}
}
