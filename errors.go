package oops

type InternalError struct {
	BaseOopsError
}

func (i *InternalError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (i InternalError) Inject(msg string, err error) Oops {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, err)
	return &i
}

type ValidationError struct {
	BaseOopsError
}

func (i *ValidationError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (i ValidationError) Inject(msg string, err error) Oops {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, err)
	return &i
}

type NotFoundError struct {
	BaseOopsError
}

func (i *NotFoundError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (n NotFoundError) Inject(msg string, err error) Oops {
	n.BaseOopsError = n.BaseOopsError.Inject(msg, err)
	return &n
}

type NotAuthorizedError struct {
	BaseOopsError
}

func (i *NotAuthorizedError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (n NotAuthorizedError) Inject(msg string, err error) Oops {
	n.BaseOopsError = n.BaseOopsError.Inject(msg, err)
	return &n
}

type TryAgainLaterError struct {
	BaseOopsError
}

func (i *TryAgainLaterError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (i TryAgainLaterError) Inject(msg string, err error) Oops {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, err)
	return &i
}

type NotAuthenticatedError struct {
	BaseOopsError
}

func (i *NotAuthenticatedError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (i NotAuthenticatedError) Inject(msg string, err error) Oops {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, err)
	return &i
}

type DeadlineExceededError struct {
	BaseOopsError
}

func (i *DeadlineExceededError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (i DeadlineExceededError) Inject(msg string, err error) Oops {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, err)
	return &i
}

type UnknownError struct {
	BaseOopsError
}

func (i *UnknownError) With(key string, value interface{}) Oops {
	i.BaseOopsError.With(key, value)
	return i
}

func (i UnknownError) Inject(msg string, err error) Oops {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, err)
	return &i
}
