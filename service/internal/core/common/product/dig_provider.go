package product

import (
	"go.uber.org/dig"

	"mall/service/internal/repository"
	"mall/service/internal/thirdparty/localcache"
	"mall/service/internal/thirdparty/mysqlcli"
)

func NewProduct(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			ProductCommon: newProductCommon(in),
		},
	}

	return self.digOut
}

type digIn struct {
	dig.In
	// 套件
	DB    mysqlcli.IMySQLClient
	Cache localcache.ILocalCache

	// Common

	// Repo
	ProductRepo repository.IProductRepo
}

type digOut struct {
	dig.Out

	ProductCommon IProductCommon
}

type packet struct {
	in digIn

	digOut
}
