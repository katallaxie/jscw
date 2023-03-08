package jscw

/*
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

type template struct {
	ctx *JSContext
}

// Set ...
func (t *template) Set(name string, val interface{}) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	switch val.(type) {
	default:
	case *FunctionTemplate:
		return errors.New("jscw: unsupported type")
	}
	runtime.KeepAlive(t)

	return nil
}
