package mysqlcli

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/dig"

	"mall/service/internal/config"
	"mall/service/internal/thirdparty/logger"
)

func NewDBClient(in digIn) IMySQLClient {
	return initWithConfig(in)
}

type digIn struct {
	dig.In

	AppConf   config.IAppConfig
	OpsConf   config.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
	AppLogger logger.ILogger `name:"appLogger"`
}
