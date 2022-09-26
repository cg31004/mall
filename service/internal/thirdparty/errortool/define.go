package errortool

import (
	"fmt"
	"sort"
)

const (
	groupCodeDB string = "00"
	errGroupLen        = 2
	errCodeLen         = 6 // groupLen(2) + "-" + codeLen(3)
)

func Define() *define {
	return &define{
		groups: newGroupRepository(),
		codes:  newCodeRepository(),
	}
}

type define struct {
	groups iGroupRepository
	codes  iCodeRepository
}

func (d *define) Group(group string) *errorGroup {
	d.groups.Add(group)
	return &errorGroup{
		codes:     d.codes,
		groups:    d.groups,
		groupCode: group,
	}
}

func (d *define) Plugin(f func(groups iGroupRepository, codes iCodeRepository) interface{}) interface{} {
	return f(d.groups, d.codes)
}

func (d *define) List() []errorString {
	keys := d.codes.Keys()
	sort.SliceStable(keys,
		func(i, j int) bool {
			return keys[i] < keys[j]
		})

	res := make([]errorString, len(keys))
	for i, v := range keys {
		if val, ok := d.codes.Get(v); ok {
			res[i] = *val
		} else {
			res[i] = errorString{}
		}
	}

	return res
}

type errorGroup struct {
	codes     iCodeRepository
	groups    iGroupRepository
	groupCode string
}

func (e *errorGroup) Error(code, message string) error {
	errCode := e.makeErrorCode(e.groups.Get(e.groupCode), code)
	err := &errorString{
		code:    errCode,
		message: message,
	}
	e.codes.Add(errCode, err)
	return err
}

func (e *errorGroup) makeErrorCode(groupCode, code string) errorCode {
	return errorCode(fmt.Sprintf("%02s-%s", groupCode, code))
}
