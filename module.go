package merr

import (
	"errors"
	"sync/atomic"

	"github.com/cornelk/hashmap"
)

type ModuleError interface {
	Name() string
	Index() uint32
	Errors() int
	New(str string) Index
	Errorf(format string, args ...any) Index
	Wrap(err error, s string) Index
	WrapIndex(err error, index Index) Index
	IndexString(index Index) string
	IndexError(index Index) error
	Is(err, target error) bool
	Has(err error) bool
	Find(err error) (Index, bool)
	Unwrap(err error) error
	As(err error, target any) bool
}

type moduleError struct {
	idx    uint32
	name   string
	count  uint32
	errors *hashmap.Map[uint64, Index]
}

func (m *moduleError) Find(err error) (Index, bool) {
	if v, ok := err.(Index); ok && m.getErrorIndex(v) != 0 {
		return v, true
	}
	return 0, false
}

func (m *moduleError) WrapIndex(err error, index Index) Index {
	return m.error(WrapIndex(err, index))
}

func (m *moduleError) error(err error) Index {
	key := hash(err.Error())
	count := atomic.AddUint32(&m.count, 1)
	idx := makeErrIndex(m.Index(), count)
	if v, ok := m.errors.GetOrInsert(key, idx); ok {
		atomic.CompareAndSwapUint32(&m.count, count, count-1)
		return v
	}
	triggerErrorHandler(err)
	globalErrors.Set(idx, err)
	return idx
}

func (m *moduleError) IndexString(index Index) string {
	return m.IndexError(index).Error()
}

func (m *moduleError) Name() string {
	return m.name
}

func (m *moduleError) Index() uint32 {
	return m.idx
}

func (m *moduleError) Errors() int {
	return m.errors.Len()
}

func (m *moduleError) New(str string) Index {
	return m.error(New(str))
}

func (m *moduleError) Wrap(err error, s string) Index {
	return m.error(WrapString(err, s))
}

func (m *moduleError) Errorf(format string, args ...any) Index {
	return m.error(Errorf(format, args...))
}

func (m *moduleError) getErrorIndex(index Index) uint32 {
	idx := (uint32(index)) ^ m.idx<<16
	if idx > m.count {
		return 0
	}
	return idx
}

func (m *moduleError) IndexError(index Index) error {
	err, ok := globalErrors.Get(index)
	if ok {
		return err
	}
	return UnknownError
}

func (m *moduleError) Is(err, target error) bool {
	return errors.Is(err, target)
}

func (m *moduleError) Unwrap(err error) error {
	return errors.Unwrap(err)
}

func (m *moduleError) As(err error, target any) bool {
	return errors.As(err, target)
}

func (m *moduleError) Has(err error) bool {
	key := hash(err.Error())
	_, ok := m.errors.Get(key)
	return ok
}

func newModuleWithIndex(name string, idx uint32) ModuleError {
	return &moduleError{
		idx:    idx,
		name:   name,
		count:  0,
		errors: hashmap.New[uint64, Index](),
	}
}

func NewModule(name string) ModuleError {
	return newModuleWithIndex(name, getModuleIndex())
}
