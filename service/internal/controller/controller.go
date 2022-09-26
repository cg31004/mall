package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"mall/service/internal/controller/handler"
	"mall/service/internal/controller/middleware"
	"mall/service/internal/core/usecase/chart"
	"mall/service/internal/core/usecase/product"
	"mall/service/internal/core/usecase/session"
	"mall/service/internal/thirdparty/logger"
)

func NewController(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			MemberCtrl:  newMember(in),
			ProductCtrl: newProduct(in),
		},
	}

	return self.digOut
}

type packet struct {
	in digIn

	digOut
}

type digIn struct {
	dig.In

	SysLogger   logger.ILogger `name:"sysLogger"`
	Request     handler.IRequestParse
	SetResponse response `optional:"true"`

	MemberIn  memberUseCaseIn
	ProductIn productUseCaseIn
}

type digOut struct {
	dig.Out

	MemberCtrl  IMemberCtrl
	ProductCtrl IProductCtrl
}

type memberUseCaseIn struct {
	dig.In

	Session session.ISessionUseCase
	Chart   chart.IMemberChartUseCase
}
type productUseCaseIn struct {
	dig.In

	Product product.IProductUseCase
}

type response struct{}

func (response) StandardResp(ctx *gin.Context, statusCode int, data interface{}) {
	middleware.SetResp(ctx, middleware.RespFormat_Standard, statusCode, "0", data)
}
