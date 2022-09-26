package session

import (
	"go.uber.org/dig"

	"simon/mall/service/internal/repository"

	"simon/mall/service/internal/thirdparty/mysqlcli"
)

func NewSession(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			SessionUseCase: newSessionUseCase(in),
		},
	}

	return self.digOut
}

type digIn struct {
	dig.In
	// 套件
	DB mysqlcli.IMySQLClient

	// Common

	// Repo
	MemberRepo  repository.IMemberRepo
	SessionRepo repository.ISessionRepo
}

type digOut struct {
	dig.Out

	SessionUseCase ISessionUseCase
}

type packet struct {
	in digIn

	digOut
}
