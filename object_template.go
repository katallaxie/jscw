package jscw

import "runtime"

// ObjectTemplate ...
type ObjectTemplate struct {
	*template
}

// NewObjectTemplate ...
func NewObjectTemplate(ctx *JSContext) *ObjectTemplate {
	tmpl := &template{}

	runtime.SetFinalizer(tmpl, (*template).finalizer)

	return &ObjectTemplate{tmpl}
}
