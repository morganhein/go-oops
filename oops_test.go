package oops

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	e := New[InternalError]("this is my message")
	i := &InternalError{}
	worked := errors.As(e, &i)
	assert.True(t, worked)
}

// TODO: This is not a great test, how should we test it?
func TestPrintNormal(t *testing.T) {
	e := Wrap[InternalError](errors.New("this is the original error"), "this is my message")
	assert.Contains(t, e.Error(), "this is my message")
	assert.Contains(t, e.Error(), "this is the original error")

	e = Wrap[InternalError](errors.New("this is the original error"), "this is my message and my name is %v", "morgan")
	assert.Contains(t, e.Error(), "morgan")
	assert.Contains(t, e.Error(), "this is my message")
	assert.Contains(t, e.Error(), "this is the original error")
	// t.Logf("%v", e)
}

func TestPrintTab(t *testing.T) {
	t.Skip("needs assertions")
	e := Wrap[InternalError](errors.New("this is the original error"), "this is my message")
	t.Logf("%t", e)
}

func TestPrintJson(t *testing.T) {
	t.Skip("needs assertions")
	e := Wrap[InternalError](errors.New("this is the original error"), "this is my message")
	t.Logf("%j", e)
}

func TestMultipleWrappedErrorsPrinting(t *testing.T) {
	t.Skip("needs assertions")
	ogErr := fmt.Errorf("puter error: %w", fmt.Errorf("middle error: %w", errors.New("internal error")))
	err := Wrap[InternalError](ogErr, "oops error")
	out, jErr := json.Marshal(err)
	assert.NoError(t, jErr)
	t.Log(string(out))
}

func TestAssertEquality(t *testing.T) {
	// write a test that asserts two of the below error of the same type are equal
	ogErr := New[NotFoundError]("this is my not found message")
	var notFoundError *NotFoundError
	equal := errors.As(ogErr, &notFoundError)
	assert.True(t, equal)

	var wrongError *InternalError
	equal = errors.As(ogErr, &wrongError)
	assert.False(t, equal)

	ogWrapped := Wrap[NotFoundError](errors.New("original not found"), "this is my not found message")
	equal = errors.As(ogWrapped, &notFoundError)
	assert.True(t, equal)
}

func TestReWrapOopsError(t *testing.T) {
	err := errors.New("this is the original error")
	ogErr := WrapInternalError(err, "this is my not found message")
	wrappedError := Wrap[NotAuthenticatedError](ogErr, "this is my not authenticated message")
	// out, jErr := json.MarshalIndent(wrappedError, "", "  ")
	// assert.NoError(t, jErr)
	// t.Log(string(out))
	t.Logf("%v", wrappedError)
	t.Logf("%j", wrappedError)
	t.Logf("%t", wrappedError)
}
