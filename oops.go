package oops

import "fmt"

type Oops interface {
	// With attaches any key/value pair to the error's metadata
	With(key string, value interface{}) Oops
	error
}

type Errors interface {
	Inject(msg string, err error) Oops
}

// New creates a new error of the specified type and returns it
func New[T Errors](format string, args ...interface{}) Oops {
	x := *(new(T))
	msg := fmt.Sprintf(format, args...)
	x2 := x.Inject(msg, nil)
	return x2
}

func Wrap[T Errors](err error, format string, args ...interface{}) Oops {
	x := *(new(T))
	msg := fmt.Sprintf(format, args...)
	x2 := x.Inject(msg, err)
	return x2
}
