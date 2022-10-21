package oops

import (
	"runtime"
)

type (
	Frame runtime.Frame
)

type InternalError struct {
	oopsError
}

func (i InternalError) inject(msg string, err error) error {
	i.oopsError = i.oopsError.Inject(msg, err)
	return &i
}

type ValidationError struct {
	oopsError
}

func (i ValidationError) inject(msg string, err error) error {
	i.oopsError = i.oopsError.Inject(msg, err)
	return &i
}

type NotFoundError struct {
	oopsError
}

func (n NotFoundError) inject(msg string, err error) error {
	n.oopsError = n.oopsError.Inject(msg, err)
	return &n
}

type NotAuthorizedError struct {
	oopsError
}

func (n NotAuthorizedError) inject(msg string, err error) error {
	n.oopsError = n.oopsError.Inject(msg, err)
	return &n
}

type TryAgainLaterError struct {
	oopsError
}

func (i TryAgainLaterError) inject(msg string, err error) error {
	i.oopsError = i.oopsError.Inject(msg, err)
	return &i
}

type NotAuthenticatedError struct {
	oopsError
}

func (i NotAuthenticatedError) inject(msg string, err error) error {
	i.oopsError = i.oopsError.Inject(msg, err)
	return &i
}

type DeadlineExceededError struct {
	oopsError
}

func (i DeadlineExceededError) inject(msg string, err error) error {
	i.oopsError = i.oopsError.Inject(msg, err)
	return &i
}

type UnknownError struct {
	oopsError
}

func (i UnknownError) inject(msg string, err error) error {
	i.oopsError = i.oopsError.Inject(msg, err)
	return &i
}
