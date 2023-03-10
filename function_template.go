package jscw

// FunctionCallback is a callback function for a JavaScript function
type FunctionCallback func(*FunctionCallbackInfo) *JSValue

// FunctionCallbackInfo is a struct that contains information about a JavaScript function
type FunctionCallbackInfo struct {
	ctx  *JSContext
	args []*jsValue
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

// FunctionTemplate ...
type FunctionTemplate struct{}

// NewFunctionTemplate ...
func NewFunctionTemplate(ctx *JSContext, cb FunctionCallback) *FunctionTemplate {
	return &FunctionTemplate{}
}
