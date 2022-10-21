package oops

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	errMsg := fmt.Sprintf("errors are bad: %v", 2)
	err := New[InternalError](errMsg)
	assert.NotNil(t, err)
	var i *InternalError
	ok := errors.As(err, &i)
	assert.True(t, ok)
	assert.Equal(t, errMsg, i.Error())
}

func TestWrap(t *testing.T) {
	innerError := errors.New("this is an inner error")
	errMsg := "this is the outer error"
	err := Wrap[InternalError](innerError, errMsg)
	assert.NotNil(t, err)
	extractedErr := errors.Unwrap(err)
	assert.Equal(t, errMsg, err.Error())
	assert.Equal(t, innerError.Error(), extractedErr.Error())
}

func TestWithMetadata(t *testing.T) {
	err := New[InternalError]("this error contains metadata").
		With("count", 1)
	assert.Error(t, err)
}

func TestAsError(t *testing.T) {
	var err error
	err = New[InternalError]("errors are bad: %v", 2)
	assert.NotNil(t, err)
}

func TestJSonFormat(t *testing.T) {
	var err error
	errMsg := fmt.Sprintf("errors are bad: %v", 2)
	err = New[InternalError](errMsg)
	writer := bytes.NewBuffer([]byte{})
	_, tErr := fmt.Fprintf(writer, "%v", err)
	assert.NoError(t, tErr)
	assert.Contains(t, writer.String(), errMsg)
}

func TestJsonFormatWith(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2).
		With("name", "morty").
		With("language", "go")
	writer := bytes.NewBuffer([]byte{})
	_, tErr := fmt.Fprintf(writer, "%V", err)
	assert.NoError(t, tErr)
	m := map[string]interface{}{}
	jErr := json.Unmarshal([]byte(writer.String()), &m)
	assert.NoError(t, jErr)
	metai, ok := m["Meta"]
	assert.True(t, ok)
	assert.NotNil(t, metai)
	meta, ok := metai.(map[string]interface{})
	assert.True(t, ok)
	assert.NotNil(t, meta)
	assert.Contains(t, meta, "name")
	assert.Contains(t, meta, "language")
	assert.Equal(t, "morty", meta["name"])
	assert.Equal(t, "go", meta["language"])
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

func TestToJson(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2).
		With("name", "morty").
		With("language", "go")
	assert.Error(t, err)
}
