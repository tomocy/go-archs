package model

import "fmt"

func isEmpty(s string) bool {
	return s == ""
}

func errorf(modelName, msg string) error {
	return fmt.Errorf("%s: %s", modelName, msg)
}
