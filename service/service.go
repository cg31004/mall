package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"simon/mall/service/internal/app/web"
	"simon/mall/service/internal/binder"
	"simon/mall/service/internal/config"
	"simon/mall/service/internal/thirdparty/logger"
	"simon/mall/service/internal/thirdparty/mysqlcli"
)

func Run() {
	setTimeZone()
	binder := binder.New()
	if err := binder.Invoke(initService); err != nil {
		panic(err)
	}

	select {}
}

type digIn struct {
	dig.In

	AppConf   config.IAppConfig
	OpsConf   config.IOpsConfig
	DB        mysqlcli.IMySQLClient
	SysLogger logger.ILogger `name:"sysLogger"`

	WebRestService web.IService
}

func initService(in digIn) {
	time.Local = time.UTC
	ctx := context.Background()

	ginMode(in)
	in.SysLogger.Info(ctx, fmt.Sprintf("[Build Info] %s", getBuildInfo()))

	srv := in.WebRestService.Run(ctx)
	serverInterrupt(ctx, in, srv)
}

func serverInterrupt(ctx context.Context, in digIn, srv *http.Server) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		select {
		case c := <-interrupt:
			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			// http close
			if err := srv.Shutdown(ctx); err != nil {
				in.SysLogger.Error(ctx, err)
			}

			// db close
			if err := in.DB.Close(); err != nil {
				in.SysLogger.Error(ctx, err)
			}

			in.SysLogger.Warn(ctx, fmt.Sprintf("Server Shutdown, osSignal: %v", c))
			os.Exit(0)
		}
	}()

}

func ginMode(in digIn) {
	gin.DisableConsoleColor()
	if in.AppConf.GetGinConfig().DebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
func setTimeZone() {
	time.Local = time.UTC
}
