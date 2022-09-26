package snowflake

import (
	"context"

	"go.uber.org/dig"

	"github.com/bwmarrin/snowflake"

	"mall/service/internal/thirdparty/logger"
)

type digIn struct {
	dig.In

	SysLogger logger.ILogger `name:"sysLogger"`
}

func NewIDGenerator(in digIn) IIDGenerator {
	node, err := snowflake.NewNode(1)
	if err != nil {
		ctx := context.Background()
		in.SysLogger.Panic(ctx, err)
		panic(err)
	}
	return &idGenerator{node: node}
}
