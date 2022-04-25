package oops

import (
	"errors"
	"fmt"
	"testing"
)

type newError struct {
	msg string
}

func (n newError) Error() string {
	return n.msg
}

func TestErrorsIs(t *testing.T) {
	sentinelErr := errors.New("this is a sentinenl error")
	genError := Wrap(sentinelErr)
	ok := errors.Is(genError, sentinelErr)
	if !ok {
		t.Fail()
	}
}

func TestErrorsAs(t *testing.T) {
	originalErr := &newError{msg: "original error"}
	genError := Wrap(originalErr)
	var detectError *newError
	ok := errors.As(genError, &detectError)
	if !ok {
		t.Fail()
	}
}

func TestWithMessage(t *testing.T) {
	originalErr := &newError{msg: "original error"}
	msgError := WithMessage("new message", originalErr)
	trace := GetTrace(msgError)
	if trace == nil {
		t.Fail()
	}
	fmt.Printf("%s\n", trace)
}

func TestTrace(t *testing.T) {
	err := New("this is a new error")
	newErr := fmt.Errorf("here's a new error: %w", err)
	eTrace := GetTrace(newErr)
	if eTrace == nil {
		t.Fail()
	}
	fmt.Printf("%s\n", eTrace)
}
