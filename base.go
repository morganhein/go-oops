package oops

import (
	"runtime"
	"strings"
)

// Actual implementation
type oopsError struct {
	actual error
	msg    string
	stack  []Frame
}

func (o *oopsError) with(key string, value interface{}) {
	panic("nothing")
}

func (o oopsError) Error() string {
	return o.msg
}

func (o oopsError) Unwrap() error {
	return o.actual
}

func (o oopsError) inject(msg string, err error) oopsError {
	// TODO: FRAMES may need to be changed, it may not go far enough back
	FRAMES := 3
	if len(strings.TrimSpace(msg)) > 0 {
		o.msg = msg
	}
	if err != nil {
		o.actual = err
	}
	// otherwise wrap and return
	pc := make([]uintptr, 15)
	n := runtime.Callers(FRAMES, pc)
	capturedFrames := runtime.CallersFrames(pc[:n])
	keepGoing := true
	var s runtime.Frame
	for keepGoing {
		s, keepGoing = capturedFrames.Next()
		// we don't want the stack trace to include anything from the protos folder on
		if strings.Contains(s.File, "/protobuf/") || strings.HasSuffix(s.File, ".pb.go") {
			keepGoing = false
			continue
		}
		if strings.Contains(s.File, "_test.go") {
			keepGoing = false
		}
		o.stack = append(o.stack, Frame(s))
	}
	return o
}
