package errors

import (
	"errors"
	"fmt"
)

func New(str string) error {
	return errors.New(str)
}

func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func UnwrapError(err error) error {
	return errors.Unwrap(err)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func WrapError(i Index, str string) error {
	return &errorErr{
		i:   i,
		str: str,
	}
}

func WrapErrorf(i Index, format string, args ...interface{}) error {
	return &errorErr{
		i:   i,
		str: fmt.Sprintf(format, args...),
		err: errors.New(format),
	}
}

func WrapErrorE(parent error, i Index, str string) error {
	return &errorErr{
		i:   i,
		str: str,
		err: parent,
	}
}
