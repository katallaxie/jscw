package jscw

/*
#include <stdlib.h>
#include <JavaScriptCore/JSBase.h>
#include <JavaScriptCore/JSContextRef.h>
*/
import "C"

type JSContext struct {
	group C.JSContextGroupRef
	ref   C.JSGlobalContextRef
}

// NewContext creates a new JSContext
func NewJSContext() *JSContext {
	ctx := new(JSContext)
	ctxGroup := C.JSContextGroupCreate()
	ctxRef := C.JSGlobalContextCreateInGroup(ctxGroup, nil)
	ctx.ref = ctxRef
	ctx.group = ctxGroup

	return ctx
}

func NewJSContextFromRef(ref C.JSGlobalContextRef) *JSContext {
	ctx := new(JSContext)
	ctx.ref = ref
	return ctx
}

func (ctx *JSContext) Convert() C.JSContextRef {
	return C.JSContextRef(ctx.ref)
}

func (ctx *JSContext) GetGlobal() *jsObject {
	return NewJSObjectFromRef(ctx.Convert(), C.JSContextGetGlobalObject(ctx.ref))
}

func (ctx *JSContext) Dispose() {
	C.JSGlobalContextRelease(ctx.ref)
	C.JSContextGroupRelease(ctx.group)
}

func (ctx *JSContext) EvaluateScript(script string, sourceUrl string) (*jsValue, error) {
	jsErr := NewJSError(ctx.Convert())
	ret := C.JSEvaluateScript(
		ctx.ref,
		NewJSString(script).ref,
		NewJSObject(ctx.Convert()).ref,
		NewJSString(sourceUrl).ref,
		C.int(0),
		&jsErr.ref)

	if jsErr.ref != nil {
		return nil, jsErr.Error(script, sourceUrl)
	}

	return NewJSValueFromRef(ctx.Convert(), ret), nil
}
