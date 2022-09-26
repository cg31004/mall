package errortool

import "errors"

type errorCode string

func Equal(err, target error) bool {
	e, ok := parse(err)
	if !ok {
		return false
	}

	return e.Is(target)
}

func Parse(err error) (*errorString, bool) {
	return parse(err)
}

func ConciseParse(err error) error {
	if e, ok := parse(err); ok {
		return e
	}

	return nil
}

func parse(err error) (*errorString, bool) {
	newError := err
	for {
		if tmp := errors.Unwrap(newError); tmp != nil {
			newError = tmp
		} else {
			break
		}
	}
	if e, ok := newError.(*errorString); ok {
		return e, true
	} else {
		return nil, false
	}
}

func New(code, message string) error {
	return &errorString{
		code:    errorCode(code),
		message: message,
	}
}

type errorString struct {
	code    errorCode
	message string
}

func (e *errorString) Error() string {
	return string(e.code) + ": " + e.message
}

func (e *errorString) GetCode() string {
	return string(e.code)
}

func (e *errorString) GetMessage() string {
	return e.message
}

func (e *errorString) Is(target error) bool {
	t, ok := parse(target)
	if !ok {
		return false
	}
	return e.GetCode() == t.GetCode()
}
