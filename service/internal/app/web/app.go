package web

import (
	"context"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"mall/service/internal/config"
	"mall/service/internal/controller"
	"mall/service/internal/controller/middleware"
	"mall/service/internal/thirdparty/errortool"
	"mall/service/internal/thirdparty/logger"
)

func NewWebRestService(in restServiceIn) IService {
	self := &restService{
		in: in,
	}

	return self
}

type restService struct {
	in restServiceIn
}

type restServiceIn struct {
	dig.In

	AppConf   config.IAppConfig
	OpsConf   config.IOpsConfig
	SysLogger logger.ILogger `name:"sysLogger"`
	AppLogger logger.ILogger `name:"appLogger"`

	//web api
	MemberCtrl controller.IMemberCtrl

	ResponseMiddleware   middleware.IResponseMiddleware
	MemberAuthMiddleware middleware.IMemberAuthMiddleware
}

type IService interface {
	Run(ctx context.Context) *http.Server
}

func (s *restService) Run(ctx context.Context) *http.Server {
	addr := s.in.AppConf.GetGinConfig().Port

	engine := s.newEngine()
	s.setRoutes(engine)

	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	s.in.SysLogger.Info(ctx, fmt.Sprintf("serve start , at => %s", addr))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			s.in.SysLogger.Info(ctx, err.Error())
		}
	}()

	return srv
}

func (s *restService) newEngine() *gin.Engine {
	return gin.New()

}

func (s *restService) setRoutes(engine *gin.Engine) {
	engine.SetTrustedProxies([]string{})

	// setting middlewares
	engine.Use(
		gin.Logger(),
		gin.Recovery(),

		s.in.ResponseMiddleware.Handle,
	)

	s.setPublicRoutes(engine)
	s.setPrivateRoutes(engine)
}

func (s *restService) setPublicRoutes(engine *gin.Engine) {
	s.setWebRoutes(engine) // Gateway 自己的功能
}

func (s *restService) setWebRoutes(engine *gin.Engine) {
	privateRouteGroup := engine.Group("mall")

	// 設定路由
	s.setApiRouters(privateRouteGroup)
}

func (s *restService) setPrivateRoutes(engine *gin.Engine) {
	privateRouteGroup := engine.Group("/_")

	// health check
	privateRouteGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// list error codes
	privateRouteGroup.GET("error-codes", getErrorCodes)

	// pprof
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	pprof.Register(engine, "/_/debug")
}

func getErrorCodes(ctx *gin.Context) {
	type code struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	codes := errortool.Codes.List()

	resp := make([]code, len(codes))
	for i, v := range codes {
		resp[i] = code{
			Code:    v.GetCode(),
			Message: v.GetMessage(),
		}
	}

	ctx.JSON(http.StatusOK, resp)
}
