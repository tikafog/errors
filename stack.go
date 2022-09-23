package errors

import (
	"fmt"
	"runtime/debug"
)

type stackError struct {
	stack []byte
	str   string
}

func (e *stackError) String() string {
	return e.Error()
}

func (e *stackError) Unwrap() error {
	return nil
}

func (e *stackError) Index() int {
	return 0
}

func (e *stackError) Message() string {
	return e.str
}

func (e *stackError) Is(target error) bool {
	if e == target {
		return true
	}
	msg, ok := target.(errMessage)
	if ok {
		return e.str == msg.Message()
	}
	return false
}

func (e *stackError) Error() string {
	return fmt.Sprintf("%v: %s", e.str, e.stack)
}

func StackError(err string) error {
	return &stackError{
		stack: debug.Stack(),
		str:   err,
	}
}
