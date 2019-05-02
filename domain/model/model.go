package model

import "fmt"
import derr "github.com/tomocy/archs/domain/error"

func isEmpty(s string) bool {
	return s == ""
}

func validationError(did, msg string) *derr.ValidationError {
	return derr.NewValidationError("failed to %s: %s", did, msg)
}

func developmentError(did, msg string) *derr.DevelopmentError {
	return derr.NewDevelopmentError("for developer: failed to %s: %s", did, msg)
}

func errorf(did, msg string) error {
	return fmt.Errorf("%s: %s", did, msg)
}
