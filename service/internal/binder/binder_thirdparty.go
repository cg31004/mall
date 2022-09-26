package binder

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/thirdparty/localcache"
	"simon/mall/service/internal/thirdparty/logger"
	"simon/mall/service/internal/thirdparty/mysqlcli"
)

func provideThirdParty(binder *dig.Container) {
	if err := binder.Provide(logger.NewAppLogger, dig.Name("appLogger")); err != nil {
		panic(err)
	}

	if err := binder.Provide(logger.NewSysLogger, dig.Name("sysLogger")); err != nil {
		panic(err)
	}

	if err := binder.Provide(mysqlcli.NewDBClient); err != nil {
		panic(err)
	}

	if err := binder.Provide(localcache.NewDefault); err != nil {
		panic(err)
	}

}
