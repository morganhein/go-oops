package oops

import (
	"errors"
)

type Oops interface {
	error
	Trace() Trace
}

func New(msg string) error {
	return &tracedError{
		original: errors.New(msg),
		trace:    addTrace(0),
	}
}

func Wrap(e error) error {
	//check if it's already a traced error
	return &tracedError{
		original: e,
		trace:    addTrace(0),
	}
}

func WithMessage(msg string) error {
	return &tracedError{
		original: errors.New(msg),
		trace:    addTrace(0),
	}
}

// GetTrace finds the earliest/first error that contains a stack trace
func GetTrace(e error) Oops {
	//iterate through each value and find any that have traces
	next := true
	var foundTraceError *tracedError
	foundTraceError, _ = e.(*tracedError)
	for next {
		unwrapped := errors.Unwrap(e)
		//find the first one to get added
		val, ok := e.(*tracedError)
		if ok {
			foundTraceError = val
		}
		if unwrapped == nil {
			next = false
			continue
		}
		e = unwrapped
	}
	//if we found a stack, return it
	if foundTraceError != nil {
		return foundTraceError
	}
	//if none found return nil
	return nil
}
