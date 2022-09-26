package binder

import (
	"mall/service/internal/app/job"
	"go.uber.org/dig"

	appWeb "mall/service/internal/app/web"
)

func provideApp(binder *dig.Container) {
	if err := binder.Provide(appWeb.NewWebRestService); err != nil {
		panic(err)
	}

	if err := binder.Provide(job.NewJobService); err != nil {
		panic(err)
	}
}
