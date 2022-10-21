![Tests](https://github.com/morganhein/go-oops/actions/workflows/tests.yml/badge.svg)

# OOPS

__Oops I did it again!__ Yet another error library

The goals of Oops are:

1. The easiest interface possible to create unique, custom errors.
2. That are fully compatible with errors.As and errors.Is for upstream detection. 
3. Automatically add stack traces.
4. Easily add custom metadata or messages.
5. Fancy tabbed printing for debugging locally, or smart JSON encoding for production logs.
6. Extendable with custom types not defined by this package.

## Usage

The simplest example:
```golang
err := oops.New[oops.InternalError]("errors can be cool too")
```
This produces an Oops error of type oops.InternalError, which can be used for upstream detection and translation.
Furthermore, the error automatically records the stack.

#### Predefined Error Types
* InternalError
* ValidationError
* NotFoundError
* NotAuthorizedError
* TryAgainLaterError
* NotAuthenticatedError
* DeadlineExceededError
* UnknownError

## Metadata
You can attach k/v pairs of information when creating an Oops Error.
```go
err := oops.New[oops.InternalError]("errors can be cool too").
	Wrap("key", interface{})
```

## Printing Verbs
This library uses several non-standard string formatting verbs for various outputs.
* `%v`: tabular, human readable format
* `%V`: tabular, human readable format with metadata
* `%s`: json, computer readable format
* `%S`: json, computer readable format with metadata

### Type aliasing
Use type aliasing if you want to rename error types:
```go
type SlowDownMcSpeedyError = TryAgainLaterError
```

### Unwrapping
This library is fully compliant with `errors.As`:
```go
 err := New[InternalError]("errors are bad: %v", 2)
 var i *InternalError
 ok := errors.As(err, &i)
 //ok is true, and you have access to the full InternalError
```

### Custom Types
To create a custom error type, you must implement the following interfaces
```go
type OopsI interface {
    With(key string, value interface{}) OopsI
}

type ErrorType interface {
    Inject(msg, errType string, err error) OopsI
}
```

The implementations __must include__ the following logic at a minimum to work with Oops:
```go
type CustomError struct {
	oops.BaseOopsError
}

func (c CustomError) Inject(msg, errType string, err error) oops.OopsI {
	c.BaseOopsError = c.BaseOopsError.Inject(msg, errType, err)
	return &c
}

func (c *CustomError) With(key string, value interface{}) oops.OopsI {
	c.BaseOopsError.With(key, value)
	return c
}
```