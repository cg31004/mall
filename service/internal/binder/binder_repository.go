package binder

import (
	"go.uber.org/dig"

	"mall/service/internal/repository"
)

func provideRepository(binder *dig.Container) {
	if err := binder.Provide(repository.NewRepository); err != nil {
		panic(err)
	}

}
