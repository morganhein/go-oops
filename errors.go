package oops

type InternalError struct {
	oopsError
}

func (i *InternalError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (i InternalError) inject(msg string, err error) Oops {
	i.oopsError = i.oopsError.inject(msg, err)
	return &i
}

type ValidationError struct {
	oopsError
}

func (i *ValidationError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (i ValidationError) inject(msg string, err error) Oops {
	i.oopsError = i.oopsError.inject(msg, err)
	return &i
}

type NotFoundError struct {
	oopsError
}

func (i *NotFoundError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (n NotFoundError) inject(msg string, err error) Oops {
	n.oopsError = n.oopsError.inject(msg, err)
	return &n
}

type NotAuthorizedError struct {
	oopsError
}

func (i *NotAuthorizedError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (n NotAuthorizedError) inject(msg string, err error) Oops {
	n.oopsError = n.oopsError.inject(msg, err)
	return &n
}

type TryAgainLaterError struct {
	oopsError
}

func (i *TryAgainLaterError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (i TryAgainLaterError) inject(msg string, err error) Oops {
	i.oopsError = i.oopsError.inject(msg, err)
	return &i
}

type NotAuthenticatedError struct {
	oopsError
}

func (i *NotAuthenticatedError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (i NotAuthenticatedError) inject(msg string, err error) Oops {
	i.oopsError = i.oopsError.inject(msg, err)
	return &i
}

type DeadlineExceededError struct {
	oopsError
}

func (i *DeadlineExceededError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (i DeadlineExceededError) inject(msg string, err error) Oops {
	i.oopsError = i.oopsError.inject(msg, err)
	return &i
}

type UnknownError struct {
	oopsError
}

func (i *UnknownError) With(key string, value interface{}) Oops {
	i.oopsError.with(key, value)
	return i
}

func (i UnknownError) inject(msg string, err error) Oops {
	i.oopsError = i.oopsError.inject(msg, err)
	return &i
}
