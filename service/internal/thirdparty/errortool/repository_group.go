package errortool

import (
	"log"
	"sync"
)

// newGroupRepository: registration error code, and check error code is unique.
func newGroupRepository() iGroupRepository {
	return &groupRepository{}
}

type iGroupRepository interface {
	Add(code string)
	Get(code string) string
}

type groupRepository struct {
	m sync.Map
}

func (c *groupRepository) Add(code string) {
	if len(code) != errGroupLen {
		log.Panicf("group error code length invalid, code: %s", code)
	}
	if _, ok := c.m.LoadOrStore(code, code); ok {
		log.Panicf("group error code duplicate definition, code: %s", code)
	}
}

func (c *groupRepository) Get(code string) string {
	val, ok := c.m.Load(code)
	if !ok {
		log.Panicf("error group code not exists, code: %s", code)
	}

	return val.(string)
}
