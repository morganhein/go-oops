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

### Type aliasing
Use type aliasing if you want to rename error types:
```go
type SlowDownMcSpeedyError = TryAgainLaterError
```