package main

import (
	"errors"
	"fmt"

	"github.com/morganhein/go-oops"
)

func main() {
	// built-in type
	err := oops.New[oops.InternalError]("errors are bad: %v", 2).
		With("name", "Morgan")
	if err == nil {
		panic("expected an error, but got nil")
	}
	var i *oops.InternalError
	ok := errors.As(err, &i)
	if !ok {
		panic("error could was not assignable to the expected type")
	}
	fmt.Printf("saw error: %v\n", i)

	// a custom error type
	err = oops.New[CustomError]("custom errors are awesome though!").
		With("name", "Morgan")
	if err == nil {
		panic("expected an error, but got nil")
	}
	var i2 *CustomError
	ok = errors.As(err, &i2)
	if !ok {
		panic("error could was not assignable to the expected type")
	}
	fmt.Printf("saw error: %v\n", i2)
}

type CustomError struct {
	oops.BaseOopsError
}

func (c CustomError) Inject(msg, errType string, err error) oops.OopsI {
	c.BaseOopsError = c.BaseOopsError.Inject(msg, errType, err)
	return &c
}

func (c *CustomError) With(key string, value interface{}) oops.OopsI {
	c.BaseOopsError.With(key, value)
	return c
}
