package oops

import (
	"fmt"
	"runtime"
)

type tracedError struct {
	original error
	trace    Trace
}

func (e *tracedError) Trace() Trace {
	return e.trace
}

func (e *tracedError) Error() string {
	return fmt.Sprintf("%v", e.original)
}

func (e *tracedError) Unwrap() error {
	if e.original != nil {
		return e.original
	}
	return nil
}

func addTrace(frames int) []Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3+frames, pc)
	capturedFrames := runtime.CallersFrames(pc[:n])
	var trace []Frame
	keepGoing := true
	var s runtime.Frame
	for keepGoing {
		s, keepGoing = capturedFrames.Next()
		trace = append(trace, Frame(s))
	}
	return trace
}
