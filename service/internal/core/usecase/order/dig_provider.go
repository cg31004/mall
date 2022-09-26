package order

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/core/common/product"
	"simon/mall/service/internal/core/common/transaction"
	"simon/mall/service/internal/repository"
	"simon/mall/service/internal/thirdparty/mysqlcli"
)

func NewOrder(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			OrderUseCase: newOrderUseCase(in),
		},
	}

	return self.digOut
}

type digIn struct {
	dig.In
	// 套件
	DB mysqlcli.IMySQLClient

	// Common
	TxnItemCommon transaction.ITxnItemCommon
	ProductCommon product.IProductCommon

	// Repo
	MemberChartRepo repository.IMemberChartRepo
	TxnRepo         repository.ITxnRepo
	TxnItemRepo     repository.ITxnItemRepo
}

type digOut struct {
	dig.Out

	OrderUseCase IOrderUseCase
}

type packet struct {
	in digIn

	digOut
}
