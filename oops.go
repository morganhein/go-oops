package oops

import (
	"errors"
	"fmt"
)

type Error interface {
	InternalError | InputError | NotFoundError | NotAuthorizedError | NotAuthenticatedError | AlreadyExistsError
	OopsError
}

type OopsError interface {
	Inject(msg string, err error) error
}

// New creates a new error of the specified type and returns it
func New[T Error](format string, values ...interface{}) error {
	x := *(new(T))
	var x2 error
	if len(values) == 0 {
		x2 = x.Inject(format, nil)
	}
	if len(values) > 0 {
		x2 = x.Inject(fmt.Sprintf(format, values...), nil)
	}
	return x2
}

func NewInternalError(format string, values ...interface{}) error {
	return New[InternalError](format, values...)
}

// Wrap creates a new error with the passed in error wrapped. If err is nil, this returns nil.
// When we wrap, detect if it's already an oops error, and if so, just augment it
func Wrap[T Error](err error, format string, values ...any) error {
	if err == nil {
		return nil
	}
	x := *(new(T))
	var x2 error
	if len(values) == 0 {
		x2 = x.Inject(format, err)
	}
	if len(values) > 0 {
		x2 = x.Inject(fmt.Sprintf(format, values...), err)
	}
	return x2
}

func WrapInternalError(err error, format string, values ...interface{}) error {
	if err == nil {
		return nil
	}
	x := InternalError{}
	var x2 error
	if len(values) == 0 {
		x2 = x.Inject(format, err)
	}
	if len(values) > 0 {
		x2 = x.Inject(fmt.Sprintf(format, values...), err)
	}
	return x2
}

func Is(err error) (OopsError, bool) {
	var internalError *InternalError
	var inputError *InputError
	var notFoundError *NotFoundError
	var notAuthorizedError *NotAuthorizedError
	var notAuthenticatedError *NotAuthenticatedError
	var alreadyExistsError *AlreadyExistsError
	switch {
	case errors.As(err, &internalError):
		return internalError, true
	case errors.As(err, &inputError):
		return inputError, true
	case errors.As(err, &notFoundError):
		return notFoundError, true
	case errors.As(err, &notAuthorizedError):
		return notAuthorizedError, true
	case errors.As(err, &notAuthenticatedError):
		return notAuthenticatedError, true
	case errors.As(err, &alreadyExistsError):
		return alreadyExistsError, true
	}

	return nil, false
}

func getStack(err error) []Frame {
	var internalError *InternalError
	var inputError *InputError
	var notFoundError *NotFoundError
	var notAuthorizedError *NotAuthorizedError
	var notAuthenticatedError *NotAuthenticatedError
	var alreadyExistsError *AlreadyExistsError
	switch {
	case errors.As(err, &internalError):
		return internalError.Stack
	case errors.As(err, &inputError):
		return inputError.Stack
	case errors.As(err, &notFoundError):
		return notFoundError.Stack
	case errors.As(err, &notAuthorizedError):
		return notAuthorizedError.Stack
	case errors.As(err, &notAuthenticatedError):
		return notAuthenticatedError.Stack
	case errors.As(err, &alreadyExistsError):
		return alreadyExistsError.Stack
	}

	return nil
}
