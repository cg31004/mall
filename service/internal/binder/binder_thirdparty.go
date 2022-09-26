package binder

import (
	"go.uber.org/dig"

	"mall/service/internal/thirdparty/localcache"
	"mall/service/internal/thirdparty/logger"
	"mall/service/internal/thirdparty/mysqlcli"
	"mall/service/internal/thirdparty/redisclient"
	"mall/service/internal/thirdparty/snowflake"
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

	if err := binder.Provide(redisclient.NewRedisClient); err != nil {
		panic(err)
	}

	if err := binder.Provide(localcache.NewDefault); err != nil {
		panic(err)
	}

	if err := binder.Provide(snowflake.NewIDGenerator); err != nil {
		panic(err)
	}

}
