package error

import "fmt"

func InInput(err error) bool {
	input, ok := err.(input)
	return ok && input.input()
}

type input interface {
	input() bool
}

func InInternal(err error) bool {
	internal, ok := err.(internal)
	return ok && internal.internal()
}

type internal interface {
	internal() bool
}

func NewValidationError(format string, a ...interface{}) *ValidationError {
	return &ValidationError{
		baseError: newBaseError(statusInput, format, a...),
	}
}

type ValidationError struct {
	*baseError
}

func NewDevelopmentError(format string, a ...interface{}) *DevelopmentError {
	return &DevelopmentError{
		baseError: newBaseError(statusInternal, format, a...),
	}
}

type DevelopmentError struct {
	*baseError
}

func newBaseError(status status, format string, a ...interface{}) *baseError {
	return &baseError{
		status: status,
		msg:    fmt.Sprintf(format, a...),
	}
}

type baseError struct {
	status status
	msg    string
}

func (e *baseError) input() bool {
	return e.status == statusInput
}

func (e *baseError) internal() bool {
	return e.status == statusInternal
}

func (e *baseError) Error() string {
	return e.msg
}

type status int

const (
	_ status = iota
	statusInput
	statusInternal
)
