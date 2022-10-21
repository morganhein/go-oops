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


## TODO:
1. Tests
   1. For `Wrap`
   2. Formatting tests that detect output
2. Taskfile
3. Pre-commit...or something
4. Finish readme
   1. Badges
   2. Examples of how to extend Oops with custom type
   3. Examples of metadata
   4. Explanation of error printing flags