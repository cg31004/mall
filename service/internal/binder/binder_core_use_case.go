package binder

import (
	"go.uber.org/dig"

	"mall/service/internal/core/usecase/chart"
	"mall/service/internal/core/usecase/product"
	"mall/service/internal/core/usecase/session"
)

func provideCoreUseCase(binder *dig.Container) {

	if err := binder.Provide(session.NewSession); err != nil {
		panic(err)
	}

	if err := binder.Provide(chart.NewChart); err != nil {
		panic(err)
	}

	if err := binder.Provide(product.NewProduct); err != nil {
		panic(err)
	}

}
