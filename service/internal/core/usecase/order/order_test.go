package order

import (
	"context"

	"github.com/stretchr/testify/suite"
)

type orderSuite struct {
	suite.Suite

	ctx context.Context
	*orderUseCase
}
