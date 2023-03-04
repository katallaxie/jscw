package jscw

// #include <stdlib.h>
// #include <string.h>
// #include <stdio.h>
// #include <jscw.h>
import "C"

// Version ...
func Version() string {
	return C.GoString(C.Version())
}
