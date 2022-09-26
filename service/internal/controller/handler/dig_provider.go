package handler

import (
	"go.uber.org/dig"

	"mall/service/internal/config"
	"mall/service/internal/thirdparty/logger"
)

func NewRequestParse(in digIn) IRequestParse {
	return &requestParseHandler{
		in: in,
	}
}

type digIn struct {
	dig.In

	AppConf   config.IAppConfig
	OpsConf   config.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
}
