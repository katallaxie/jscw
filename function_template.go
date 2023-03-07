package jscw

// #include <stdlib.h>
// #include <JavaScriptCore/JSBase.h>
// #include <JavaScriptCore/JSContextRef.h>
import "C"

// FunctionCallback is a callback function for a JavaScript function
type FunctionCallback func(*FunctionCallbackInfo) *JSValue

// FunctionCallbackInfo is a struct that contains information about a JavaScript function
type FunctionCallbackInfo struct {
	ctx  *JSContext
	args []*JSValue
	this *JSObject
}

// This ...
func (info *FunctionCallbackInfo) This() *JSObject {
	return info.this
}

// Context ...
func (info *FunctionCallbackInfo) Context() *JSContext {
	return info.ctx
}

// Object ...
func (info *FunctionCallbackInfo) Object() *JSObject {
	return info.this
}
