package oops

type InternalError struct {
	BaseOopsError
}

func (i *InternalError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (i InternalError) Inject(msg, errType string, err error) OopsI {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, errType, err)
	return &i
}

type ValidationError struct {
	BaseOopsError
}

func (i *ValidationError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (i ValidationError) Inject(msg, errType string, err error) OopsI {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, errType, err)
	return &i
}

type NotFoundError struct {
	BaseOopsError
}

func (i *NotFoundError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (n NotFoundError) Inject(msg, errType string, err error) OopsI {
	n.BaseOopsError = n.BaseOopsError.Inject(msg, errType, err)
	return &n
}

type NotAuthorizedError struct {
	BaseOopsError
}

func (i *NotAuthorizedError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (n NotAuthorizedError) Inject(msg, errType string, err error) OopsI {
	n.BaseOopsError = n.BaseOopsError.Inject(msg, errType, err)
	return &n
}

type TryAgainLaterError struct {
	BaseOopsError
}

func (i *TryAgainLaterError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (i TryAgainLaterError) Inject(msg, errType string, err error) OopsI {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, errType, err)
	return &i
}

type NotAuthenticatedError struct {
	BaseOopsError
}

func (i *NotAuthenticatedError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (i NotAuthenticatedError) Inject(msg, errType string, err error) OopsI {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, errType, err)
	return &i
}

type DeadlineExceededError struct {
	BaseOopsError
}

func (i *DeadlineExceededError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (i DeadlineExceededError) Inject(msg, errType string, err error) OopsI {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, errType, err)
	return &i
}

type UnknownError struct {
	BaseOopsError
}

func (i *UnknownError) With(key string, value interface{}) OopsI {
	i.BaseOopsError.With(key, value)
	return i
}

func (i UnknownError) Inject(msg, errType string, err error) OopsI {
	i.BaseOopsError = i.BaseOopsError.Inject(msg, errType, err)
	return &i
}
