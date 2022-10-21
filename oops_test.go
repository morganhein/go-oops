package oops

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	err := New[InternalError]("errors are bad: %v", 2)
	assert.NotNil(t, err)
	var i *InternalError
	ok := errors.As(err, &i)
	assert.True(t, ok)
	assert.Equal(t, "errors are bad: 2", i.Error())
}
