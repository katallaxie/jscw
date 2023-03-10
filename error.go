package jscw

/*
#include <stdlib.h>
#include <JavaScriptCore/JSBase.h>
#include <JavaScriptCore/JSContextRef.h>
#include <JavaScriptCore/JSValueRef.h>
*/
import "C"

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type JSError struct {
	ctx C.JSContextRef
	ref C.JSValueRef
}

// NewJSError ...
func NewJSError(ctx C.JSContextRef) *JSError {
	err := new(JSError)
	err.ctx = ctx

	return err
}

func NewJSErrorFromRef(ctx C.JSContextRef, ref C.JSValueRef) *JSError {
	err := new(JSError)
	err.ctx = ctx
	err.ref = ref
	return err
}

func (e *JSError) Value() *jsValue {
	return NewJSValueFromRef(e.ctx, e.ref)
}

func (e *JSError) Object() *jsObject {
	return e.Value().Object()
}

func (e *JSError) Error(script string, sourceUrl string) error {
	var errSb bytes.Buffer
	errObj := e.Object()

	if len(sourceUrl) > 0 {
		errSb.WriteString(sourceUrl)
	} else {
		errSb.WriteString("[Unknown Url]")
	}
	errSb.WriteRune(':')
	line := errObj.GetProperty("line").String()
	lineNum, _ := strconv.Atoi(line)
	errSb.WriteString(line)
	errSb.WriteRune('\n')
	if len(script) > 0 {
		errSb.WriteString(strings.Split(script, "\n")[lineNum-1])
		errSb.WriteRune('\n')
		errSb.WriteString("^^^")
	} else {
		errSb.WriteString("[Unknown Script]")
	}
	errSb.WriteRune('\n')
	errSb.WriteString(errObj.GetProperty("name").String())
	errSb.WriteString(": ")
	errSb.WriteString(errObj.GetProperty("message").String())

	return errors.New(errSb.String())
}
