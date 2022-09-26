package errortool

import (
	"log"
	"sync"
)

// newCodeRepository: registration error code, and check error code is unique.
func newCodeRepository() iCodeRepository {
	return &codeRepository{}
}

type iCodeRepository interface {
	Add(code errorCode, err *errorString)
	Get(code errorCode) (*errorString, bool)
	Keys() []errorCode
}

type codeRepository struct {
	m sync.Map
}

func (c *codeRepository) Add(code errorCode, err *errorString) {
	if len(code) != errCodeLen {
		log.Panicf("error code length invalid, code: %s", code)
	}

	if _, ok := c.m.LoadOrStore(code, err); ok {
		log.Panicf("error code duplicate definition, code: %s", code)
	}
}

func (c *codeRepository) Get(code errorCode) (*errorString, bool) {
	val, ok := c.m.Load(code)
	return val.(*errorString), ok
}

func (c *codeRepository) Keys() []errorCode {
	result := make([]errorCode, 0)

	c.m.Range(func(key, value interface{}) bool {
		result = append(result, key.(errorCode))
		return true
	})

	return result
}
