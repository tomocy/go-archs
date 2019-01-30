package usecase

type NoSuchUserError interface {
	error
	noSuchUserError()
}

func newNoSuchUserError() NoSuchUserError {
	return new(noSuchUserError)
}

func IsNoSuchUserError(err error) bool {
	_, ok := err.(NoSuchUserError)
	return ok
}

type noSuchUserError struct {
}

func (e noSuchUserError) Error() string {
	return "no such user"
}

func (e noSuchUserError) noSuchUserError() {
}
