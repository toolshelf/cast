package internal

import (
	"errors"
	"fmt"
)

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

// ErrSafe indicates that a value is out of safe integer range [-(1<<53-1), 1<<53 -1].
var ErrSafe = errors.New("value isn't a safe integer")

type CastError struct {
	In      interface{}
	OutType string
	Err     error
}

func (e *CastError) Error() string {
	return fmt.Errorf("unable to cast %#v of type %T to %s: %s", e.In, e.In, e.OutType, e.Err.Error()).Error()
}

func (e *CastError) Unwrap() error { return e.Err }

func RangeError(in interface{}, outType string) *CastError {
	return &CastError{in, outType, ErrRange}
}

func SyntaxError(in interface{}, outType string) *CastError {
	return &CastError{in, outType, ErrSyntax}
}

func SafeError(in interface{}, outType string) *CastError {
	return &CastError{in, outType, ErrSafe}
}
