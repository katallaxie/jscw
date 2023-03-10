package jscw

/*
#include <stdlib.h>
#include <JavaScriptCore/JSBase.h>
#include <JavaScriptCore/JSContextRef.h>
#include <JavaScriptCore/JSValueRef.h>
#include <JavaScriptCore/JSObjectRef.h>
#include <JavaScriptCore/JSStringRef.h>
*/
import "C"
import "fmt"

// JSValue ...
type JSValue interface {
	fmt.Stringer
}

type jsValue struct {
	ref C.JSValueRef
	ctx C.JSContextRef

	JSValue
	JSObject
}

// NewJSValueFromRef ...
func NewJSValueFromRef(ctx C.JSContextRef, ref C.JSValueRef) *jsValue {
	jsVal := new(jsValue)
	jsVal.ctx = ctx
	jsVal.ref = ref
	return jsVal
}

func NewJSUndefined(ctx C.JSContextRef) *jsValue {
	jsVal := new(jsValue)
	jsVal.ctx = ctx
	jsVal.ref = C.JSValueMakeUndefined(ctx)
	return jsVal
}

func NewJSUndefinedFromCtxRef(ref C.JSContextRef) *jsValue {
	jsVal := new(jsValue)
	jsVal.ref = C.JSValueMakeUndefined(ref)
	return jsVal
}

// String ...
func (v *jsValue) String() string {
	jsErr := NewJSError(v.ctx)
	jsStr := NewJSStringFromRef(C.JSValueToStringCopy(v.ctx, v.ref, &jsErr.ref))
	defer jsStr.Dispose()
	return jsStr.String()
}

// IsString ...
func (v *jsValue) IsString() bool {
	return false
}

func (v *jsValue) Object() *jsObject {
	jsErr := NewJSError(v.ctx)
	ret := C.JSValueToObject(v.ctx, v.ref, &jsErr.ref)
	if jsErr.ref != nil {
		panic("js err")
	}
	return NewJSObjectFromRef(v.ctx, ret)
}

func (v *jsValue) Float64() float64 {
	jsErr := NewJSError(v.ctx)
	ret := C.JSValueToNumber(v.ctx, v.ref, &jsErr.ref)
	if jsErr.ref != nil {
		panic("js err")
	}
	return float64(ret)
}

func (v *jsValue) Uint8() uint8 {
	jsErr := NewJSError(v.ctx)
	ret := C.JSValueToNumber(v.ctx, v.ref, &jsErr.ref)
	if jsErr.ref != nil {
		panic("js err")
	}
	return uint8(ret)
}
