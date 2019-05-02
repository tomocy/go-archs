package error

import "fmt"

func InUserRegistration(err error) bool {
	ur, ok := err.(userRegistration)
	return ok && ur.userRegistration()
}

type userRegistration interface {
	userRegistration() bool
}

func InInput(err error) bool {
	input, ok := err.(input)
	return ok && input.input()
}

type input interface {
	input() bool
}

func NewValidationError(kind Kind, format string, a ...interface{}) *ValidationError {
	return &ValidationError{
		usecaseError: newUsecaseError(kind, statusInput, format, a...),
	}
}

type ValidationError struct {
	*usecaseError
}

func newUsecaseError(kind Kind, status status, format string, a ...interface{}) *usecaseError {
	return &usecaseError{
		kind:   kind,
		status: status,
		msg:    fmt.Sprintf(format, a...),
	}
}

type usecaseError struct {
	kind   Kind
	status status
	msg    string
}

type Kind int

const (
	_ Kind = iota
	KindUserRegistration
)

type status int

const (
	_ status = iota
	statusInput
)

func (e *usecaseError) userRegistration() bool {
	return e.kind == KindUserRegistration
}

func (e *usecaseError) input() bool {
	return e.status == statusInput
}

func (e *usecaseError) Error() string {
	return e.msg
}
