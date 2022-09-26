package snowflake

import "github.com/bwmarrin/snowflake"

//go:generate mockery --name IIDGenerator --structname MockIDGenerator --output mock_snow --outpkg mock_snow --filename mock_logger.go --with-expecter

type IIDGenerator interface {
	GenerateInt64() int64
}

type idGenerator struct {
	node *snowflake.Node
}

func (id *idGenerator) GenerateInt64() int64 {
	return id.node.Generate().Int64()
}
