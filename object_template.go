package jscw

// #include <stdlib.h>
// #include <JavaScriptCore/JSBase.h>
// #include <JavaScriptCore/JSContextRef.h>
import "C"

// ObjectTemplate ...
type ObjectTemplate struct{}

// NewObjectTemplate ...
func NewObjectTemplate(ctx *JSContext) *ObjectTemplate {
	return &ObjectTemplate{}
}
