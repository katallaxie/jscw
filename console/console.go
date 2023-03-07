package console

import (
	"io"
	"os"
)

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
func (c *console) GetLogFunctionCallback() {}
