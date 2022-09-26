package binder

import (
	"go.uber.org/dig"

	"mall/service/internal/controller"
	"mall/service/internal/controller/handler"
	"mall/service/internal/controller/middleware"
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
