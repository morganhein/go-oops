package main

import (
	"errors"

	"github.com/morganhein/go-oops"
)

func main() {
	err := oops.New[oops.InternalError]("errors are bad: %v", 2).
		With("name", "Morgan")
	if err != nil {
		panic(err)
	}
	var i *oops.InternalError
	ok := errors.As(err, &i)
	if !ok {
		panic("error could was not assignable to the expected type")
	}
}
