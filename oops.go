package oops

import "fmt"

type OopsI interface {
	// With attaches any key/value pair to the error's metadata
	With(key string, value interface{}) OopsI
	error
}

type ErrorType interface {
	// Inject is used to fill out a BaseError for any new OopsError
	// msg: error message for the new error
	// errType: the stringified name of the new error
	// err: the original error to wrap, if any
	Inject(msg, errType string, err error) OopsI
}

// New creates a new error of the specified type and returns it
func New[T ErrorType](format string, args ...interface{}) OopsI {
	newt := *(new(T))
	t := getType(newt)
	msg := fmt.Sprintf(format, args...)
	e := newt.Inject(msg, t, nil)
	return e
}

func Wrap[T ErrorType](err error, format string, args ...interface{}) OopsI {
	newt := *(new(T))
	t := getType(newt)
	msg := fmt.Sprintf(format, args...)
	e := newt.Inject(msg, t, err)
	return e
}
