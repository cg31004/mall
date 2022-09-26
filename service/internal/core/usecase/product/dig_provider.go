package product

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/repository"
	"simon/mall/service/internal/thirdparty/mysqlcli"
)

func NewProduct(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			ProductUseCase: newProductUseCase(in),
		},
	}

	return self.digOut
}

type digIn struct {
	dig.In
	// 套件
	DB mysqlcli.IMySQLClient

	// Common

	// Repo
	ProductRepo repository.IProductRepo
}

type digOut struct {
	dig.Out

	ProductUseCase IProductUseCase
}

type packet struct {
	in digIn

	digOut
}
