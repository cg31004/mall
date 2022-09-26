package redisclient

import (
	"go.uber.org/dig"

	"mall/service/internal/config"
)

func NewRedisClient(in digIn) IRedisClient {
	return initWithConfig(in)
}

type digIn struct {
	dig.In

	OpsConf config.IOpsConfig
}

func NewMockClient() *RedisClient{
	return self
}
