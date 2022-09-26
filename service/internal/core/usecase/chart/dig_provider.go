package chart

import (
	"go.uber.org/dig"

	"mall/service/internal/core/common/product"
	"mall/service/internal/repository"
	"mall/service/internal/thirdparty/mysqlcli"
)

func NewChart(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			MemberChartUseCase: newMemberChartUseCase(in),
		},
	}

	return self.digOut
}

type digIn struct {
	dig.In
	// 套件
	DB mysqlcli.IMySQLClient

	// Common
	ProductCommon product.IProductCommon

	// Repo
	MemberChartRepo repository.IMemberChartRepo
}

type digOut struct {
	dig.Out

	MemberChartUseCase IMemberChartUseCase
}

type packet struct {
	in digIn

	digOut
}
