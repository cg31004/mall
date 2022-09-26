package middleware

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/core/usecase/session"
	"simon/mall/service/internal/thirdparty/logger"
)

func NewResponseMiddleware(in digIn) IResponseMiddleware {
	resp := &responseMiddleware{
		in: in,
	}

	return resp
}

func NewUserAuthMiddleware(in digIn) IMemberAuthMiddleware {
	return &userAuthMiddleware{
		in: in,
	}
}

type digIn struct {
	dig.In

	SysLogger logger.ILogger `name:"sysLogger"`
	AppLogger logger.ILogger `name:"appLogger"`

	SessionUseCase session.ISessionUseCase
}
