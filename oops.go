package oops

import (
	"errors"
	"fmt"
	"runtime"
)

type Oops interface {
	Trace() *trace
	error
}

func Wrap[T error](e T) *Error[T] {
	//detect if T is not an error using reflection?

	return &Error[T]{
		trace:    getStack(0),
		original: e,
	}
}

func getStack(frames int) []Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3+frames, pc)
	capturedFrames := runtime.CallersFrames(pc[:n])
	trace := []Frame{}
	keepGoing := true
	var s runtime.Frame
	for keepGoing {
		s, keepGoing = capturedFrames.Next()
		trace = append(trace, Frame(s))
	}
	return trace
}

type Error[T error] struct {
	msg      string
	original error
	trace    trace
}

func (e *Error[T]) Error() string {
	return fmt.Sprintf("%v", e.original)
}

func (e *Error[T]) Is(target error) bool {
	return errors.Is(e.original, target)
}

func (e *Error[T]) As(target any) bool {
	return errors.As(e.original, target)
}
