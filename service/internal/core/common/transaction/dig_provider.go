package transaction

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/repository"
	"simon/mall/service/internal/thirdparty/localcache"
	"simon/mall/service/internal/thirdparty/mysqlcli"
)

func NewTxn(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			TxnItemCommon: newTxnItemCommon(in),
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
	TxnItemRepo repository.ITxnItemRepo
}

type digOut struct {
	dig.Out

	TxnItemCommon ITxnItemCommon
}

type packet struct {
	in digIn

	digOut
}
