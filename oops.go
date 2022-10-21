package oops

import "fmt"

type OopsI interface {
	// With attaches any key/value pair to the error's metadata
	With(key string, value interface{}) OopsI
	error
}

type ErrorType interface {
	Inject(msg string, err error) OopsI
}

// New creates a new error of the specified type and returns it
func New[T ErrorType](format string, args ...interface{}) OopsI {
	newt := *(new(T))
	msg := fmt.Sprintf(format, args...)
	e := newt.Inject(msg, nil)
	return e
}

func Wrap[T ErrorType](err error, format string, args ...interface{}) OopsI {
	newt := *(new(T))
	msg := fmt.Sprintf(format, args...)
	e := newt.Inject(msg, err)
	return e
}
