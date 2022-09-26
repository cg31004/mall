package binder

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/controller"
	"simon/mall/service/internal/controller/handler"
	"simon/mall/service/internal/controller/middleware"
)

func provideController(binder *dig.Container) {
	if err := binder.Provide(handler.NewRequestParse); err != nil {
		panic(err)
	}

	if err := binder.Provide(middleware.NewResponseMiddleware); err != nil {
		panic(err)
	}

	if err := binder.Provide(middleware.NewUserAuthMiddleware); err != nil {
		panic(err)
	}

	if err := binder.Provide(controller.NewController); err != nil {
		panic(err)
	}

}
