package repository

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/config"
	"simon/mall/service/internal/thirdparty/localcache"
)

func NewRepository(in repositoryIn) repositoryOut {
	self := &repository{
		in: in,

		repositoryOut: repositoryOut{
			MemberRepo:      newMemberRepo(in),
			MemberChartRepo: newMemberChartRepo(in),
			SessionRepo:     newSessionRepoByRedis(in),
			TxnItemRepo:     newTxnItemRepo(in),
			TxnRepo:         newTxnRepo(in),
			ProductRepo:     newProductRepo(in),
		},
	}

	return self.repositoryOut
}

type repositoryIn struct {
	dig.In

	LocalCache localcache.ILocalCache
	AppConf    config.IAppConfig
}

type repository struct {
	in repositoryIn

	repositoryOut
}

type repositoryOut struct {
	dig.Out

	MemberRepo      IMemberRepo
	SessionRepo     ISessionRepo
	TxnItemRepo     ITxnItemRepo
	TxnRepo         ITxnRepo
	MemberChartRepo IMemberChartRepo
	ProductRepo     IProductRepo
}
