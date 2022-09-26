package binder

import (
	"sync"

	"go.uber.org/dig"
)

var (
	binder *dig.Container
	once   sync.Once
)

func New() *dig.Container {
	once.Do(func() {
		binder = dig.New()

		provideThirdParty(binder)
		provideConfig(binder)
		provideApp(binder)
		provideController(binder)
		provideCoreUseCase(binder)
		provideCoreCommon(binder)
		provideRepository(binder)
	})

	return binder
}
