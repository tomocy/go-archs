package validator

import "fmt"

type EmptyError interface {
	error
	emptyError()
}

func newEmptyError(fieldName string) EmptyError {
	return &emptyError{
		fieldName: fieldName,
	}
}

func IsEmptyError(err error) bool {
	_, ok := err.(EmptyError)
	return ok
}

type emptyError struct {
	fieldName string
}

func (e emptyError) Error() string {
	return fmt.Sprintf("%s should not be empty", e.fieldName)
}

func (e emptyError) emptyError() {
}
