package oops

import "fmt"

type OopsI interface {
	// With attaches any key/value pair to the error's metadata
	With(key string, value interface{}) OopsI
	error
}

type Injectable interface {
	Inject(msg string, err error) OopsI
}

// New creates a new error of the specified type and returns it
func New[T Injectable](format string, args ...interface{}) OopsI {
	x := *(new(T))
	msg := fmt.Sprintf(format, args...)
	x2 := x.Inject(msg, nil)
	return x2
}

func Wrap[T Injectable](err error, format string, args ...interface{}) OopsI {
	x := *(new(T))
	msg := fmt.Sprintf(format, args...)
	x2 := x.Inject(msg, err)
	return x2
}
