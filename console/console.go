package console

import (
	"io"
	"os"

	"github.com/katallaxie/jscw"
)

// #include <stdlib.h>
// #include <JavaScriptCore/JSBase.h>
// #include <JavaScriptCore/JSContextRef.h>
import "C"

type console struct {
	out        io.Writer
	methodName string
}

// New creates a new console object
func New() *console {
	c := new(console)
	c.out = os.Stdout
	c.methodName = "log"

	return c
}

// GetLogFunctionCallback returns a callback function for the console.log method
func (c *console) GetLogFunctionCallback() jscw.FunctionCallback {
	return func(info *jscw.FunctionCallbackInfo) *jscw.JSValue {
		return nil
	}
}

// Inject ...
func (c *console) Inject(ctx *jscw.JSContext) {

}
