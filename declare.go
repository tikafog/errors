package merr

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/cornelk/hashmap"
)

// Index is used for make error index
type Index uint32

var (
	moduleIndex         uint32
	modulesMu           sync.RWMutex
	moduleNames         []uint64
	moduleErrors        = hashmap.New[uint64, ModuleError]()
	globalErrors        = hashmap.New[Index, error]()
	triggerErrorHandler = func(err error) {}
)

var (
	UnknownError  = IndexError(0)
	UnknownModule = registerModuleWithIndex("unknown", 0)
)

func IndexNew(str string) Index {
	return UnknownModule.New(str)
}

func IndexErrorf(format string, args ...interface{}) error {
	return UnknownModule.Errorf(format, args...)
}

func getModuleIndex() uint32 {
	modulesMu.Lock()
	moduleNames = append(moduleNames, 0)
	modulesMu.Unlock()
	return atomic.AddUint32(&moduleIndex, 1)
}

func Module(name string) ModuleError {
	key := hash(name)
	e, ok := moduleErrors.Get(key)
	if ok {
		return e
	}
	return UnknownModule
}

func moduleNameHash(index Index) uint64 {
	idx := index.ModuleIndex()
	if idx < uint32(len(moduleNames)) {
		return atomic.LoadUint64(&moduleNames[idx])

	}
	return hash(UnknownModule.Name())
}

func ModuleFromIndex(index Index) ModuleError {
	key := moduleNameHash(index)
	if v, ok := moduleErrors.Get(key); ok {
		return v
	}
	return UnknownModule
}

func registerModuleWithIndex(name string, idx uint32) ModuleError {
	key := hash(strings.ToLower(name))
	m := newModuleWithIndex(name, idx)
	if v, ok := moduleErrors.GetOrInsert(key, m); ok {
		return v
	}
	atomic.StoreUint64(&moduleNames[m.Index()], key)
	return m
}

func RegisterModule(name string) ModuleError {
	v, ok := moduleErrors.Get(hash(name))
	if ok {
		return v
	}
	return registerModuleWithIndex(name, getModuleIndex())
}

func RegisterErrorHandler(fn func(err error)) {
	triggerErrorHandler = fn
}

// Index ...
func (e Index) Index() Index {
	return e
}

// ModuleName ...
func (e Index) ModuleName() string {
	return ModuleFromIndex(e).Name()
}

// ModuleIndex ...
func (e Index) ModuleIndex() uint32 {
	return uint32(e) >> 16
}

func (e Index) Error() string {
	return e.String()
}

// String gets the string value of Index
func (e Index) String() string {
	if v, ok := globalErrors.Get(e); ok {
		return fmt.Sprintf("Module[%v]: %v", e.ModuleName(), v.Error())
	}
	return "unknown error"
	//err := Module(e.moduleNameHash()).IndexError(e)

	//return
}

// IndexModule ...
func IndexModule(e uint32) uint32 {
	return e << 16
}

// ModuleIndex ...
//func moduleIndex(e Index) uint32 {
//	return uint32(e) >> 16
//}

func (e Index) Module() ModuleError {
	return Module(e.ModuleName())
}
