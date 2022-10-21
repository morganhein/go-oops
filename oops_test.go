package oops

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2).
		With("name", "morty").
		With("language", "go")
	assert.NotNil(t, err)
	var i *InternalError
	ok := errors.As(err, &i)
	assert.True(t, ok)
	assert.Equal(t, "errors are bad: 2", i.Error())
}

func TestWrap(t *testing.T) {
	innerError := errors.New("this is an inner error")
	err := Wrap[InternalError](innerError, "this is the outer error")
	assert.NotNil(t, err)
	t.Logf("\n%s", err)
}

func TestAsError(t *testing.T) {
	var err error
	err = New[InternalError]("errors are bad: %v", 2)
	assert.NotNil(t, err)
}

func TestJSonFormat(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2)
	t.Logf("\n%v", err)
}

func TestJsonFormatWith(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2).
		With("name", "morty").
		With("language", "go")
	t.Logf("\n%V", err)
}

func TestTabFormat(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2)
	t.Logf("\n%s", err)
}

func TestTabFormatWith(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2).
		With("name", "morty").
		With("language", "go")
	t.Logf("\n%S", err)
}
