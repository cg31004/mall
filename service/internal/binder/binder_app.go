package binder

import (
	"go.uber.org/dig"

	appWeb "simon/mall/service/internal/app/web"
)

func provideApp(binder *dig.Container) {
	if err := binder.Provide(appWeb.NewWebRestService); err != nil {
		panic(err)
	}

}
