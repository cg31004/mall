package binder

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/core/common/product"
	"simon/mall/service/internal/core/common/transaction"
)

func provideCoreCommon(binder *dig.Container) {

	if err := binder.Provide(product.NewProduct); err != nil {
		panic(err)
	}

	if err := binder.Provide(transaction.NewTxn); err != nil {
		panic(err)
	}
}
