package binder

import (
	"go.uber.org/dig"

	"mall/service/internal/core/common/product"
)

func provideCoreCommon(binder *dig.Container) {

	if err := binder.Provide(product.NewProduct); err != nil {
		panic(err)
	}
}
