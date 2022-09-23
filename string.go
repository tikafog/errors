package merr

import (
	"fmt"
)

type stringErr struct {
	s   string
	err error
}

func (e *stringErr) Error() string {
	return e.String()
}

func (e *stringErr) String() string {
	if e.err != nil {
		return fmt.Sprintf("%v: %v", e.s, e.err)
	}
	return e.s
}

func (e *stringErr) Message() string {
	return e.String()
}

func (e *stringErr) Unwrap() error {
	return e.err
}

func (e *stringErr) Is(target error) bool {
	if e == target {
		return true
	}
	idx, ok := target.(errMessage)
	if ok && e.Message() == idx.Message() {
		return true
	}
	return false
}

// StringError
// @param string
// @return error
// Decrypted use ErrorString instead
func StringError(s string) error {
	return &stringErr{s: s}
}

func ErrorString(s string) error {
	return &stringErr{s: s}
}

func WrapString(err error, str string) error {
	if err == nil {
		return &stringErr{s: str}
	}
	return &stringErr{s: str, err: err}
}

func WrapStringN(err error, str string) error {
	if err == nil {
		return nil
	}
	return &stringErr{s: str, err: err}
}
