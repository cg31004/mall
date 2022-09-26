package repository

import (
	"go.uber.org/dig"

	"mall/service/internal/config"
	"mall/service/internal/thirdparty/localcache"
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
			PaymentRepo:     newPaymentRepo(in),
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
	PaymentRepo     IPaymentRepo
	ProductRepo     IProductRepo
}
