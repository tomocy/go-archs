package model

import "fmt"

func isEmpty(s string) bool {
	return s == ""
}

func errorf(did, msg string) error {
	return fmt.Errorf("%s: %s", did, msg)
}
