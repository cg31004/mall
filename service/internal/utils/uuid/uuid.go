package uuid

import "github.com/google/uuid"

//go:generate mockery --name IUuid --structname MockUuid --output mock_uuid --outpkg mock_uuid --filename mock_uuid.go --with-expecter

type IUuid interface {
	GetUUID() string
}

type idGenerator struct{}

func NewIdGenerator() IUuid {
	return &idGenerator{}
}

func (*idGenerator) GetUUID() string {
	return uuid.New().String()
}
