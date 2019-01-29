package usecase

import "fmt"

type DuplicatedEmailError interface {
	error
	duplicatedEmail() string
}

type duplicatedEmailError struct {
	email string
}

func newDuplicatedEmailError(email string) DuplicatedEmailError {
	return &duplicatedEmailError{
		email: email,
	}
}

func IsDuplicatedEmailError(err error) bool {
	_, ok := err.(DuplicatedEmailError)
	return ok
}

func (e duplicatedEmailError) Error() string {
	return fmt.Sprintf("duplicated email: %s\n", e.email)
}

func (e duplicatedEmailError) duplicatedEmail() string {
	return e.email
}

type IncorrectCredentialError interface {
	error
	incorrectCredentialError()
}

type incorrectCredentialError struct {
}

func newIncorrectCredentialError() IncorrectCredentialError {
	return new(incorrectCredentialError)
}

func IsIncorrectCredentialError(err error) bool {
	_, ok := err.(IncorrectCredentialError)
	return ok
}

func (e incorrectCredentialError) Error() string {
	return "incorrect credential error"
}

func (e incorrectCredentialError) incorrectCredentialError() {
}
