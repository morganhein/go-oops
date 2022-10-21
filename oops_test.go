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

// TODO: this test could be better.... checking for the metadata seems an obvious choice
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
	var m map[string]interface{}
	mErr := json.Unmarshal([]byte(writer.String()), &m)
	assert.NoError(t, mErr)
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
	var err error
	errMsg := fmt.Sprintf("errors are bad: %v", 2)
	err = New[InternalError](errMsg)
	writer := bytes.NewBuffer([]byte{})
	_, tErr := fmt.Fprintf(writer, "%s", err)
	assert.NoError(t, tErr)
	assert.Contains(t, writer.String(), errMsg)
}

func TestTabFormatWith(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2).
		With("name", "morty").
		With("language", "go")
	writer := bytes.NewBuffer([]byte{})
	_, tErr := fmt.Fprintf(writer, "%S", err)
	assert.NoError(t, tErr)
	assert.Contains(t, writer.String(), "language")
	assert.Contains(t, writer.String(), "go")
	assert.Contains(t, writer.String(), "name")
	assert.Contains(t, writer.String(), "morty")
}

func TestToJson(t *testing.T) {
	var err error
	err = New[InternalError]("errors are bad: %v", 2).
		With("name", "morty").
		With("language", "go")
	d, err := json.Marshal(err)
	assert.NoError(t, err)
	var m map[string]interface{}
	jErr := json.Unmarshal(d, &m)
	assert.NoError(t, jErr)
	ty, ok := m["Type"]
	assert.True(t, ok)
	assert.Equal(t, "InternalError", ty)
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
