package binder

import (
	"go.uber.org/dig"

	"mall/service/internal/config"
)

func provideConfig(binder *dig.Container) {
	if err := binder.Provide(config.NewAppConfig); err != nil {
		panic(err)
	}

	if err := binder.Provide(config.NewOpsConfig); err != nil {
		panic(err)
	}
}
