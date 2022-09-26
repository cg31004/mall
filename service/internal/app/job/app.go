package job

import (
	"go.uber.org/dig"

	"mall/service/internal/config"
	"mall/service/internal/controller"
	"mall/service/internal/thirdparty/cron"
	"mall/service/internal/thirdparty/logger"
)

func NewJobService(in jobServiceIn) IService {
	self := &jobService{
		in: in,
	}

	return self
}

type jobServiceIn struct {
	dig.In

	AppConf   config.IAppConfig
	OpsConf   config.IOpsConfig
	AppLogger logger.ILogger `name:"appLogger"`
	SysLogger logger.ILogger `name:"sysLogger"`
	Cron      cron.ICronJob

	BankCtrl controller.IBankCtrl
}

type IService interface {
	Run()
}

type jobService struct {
	in jobServiceIn
}

func (s *jobService) Run() {
	s.schedule()
	go s.in.Cron.Start()
}

func (s *jobService) schedule() {
	//s.in.Cron.AddFunc("0 * * * * *", s.in.ExampleCtrl.Get)
	//s.in.Cron.AddScheduleFunc(2*time.Second, s.in.ExampleCtrl.Get)
	//s.in.Cron.AddScheduleFunc(config.NewAppConfig().GetCronConfig().PointStatistcalPeriodSec*time.Second, s.in.PointStatistical.PointStatistical)
	s.in.Cron.AddFunc("0 0 4 * * *", s.in.BankCtrl.SyncBankInfo)
}
