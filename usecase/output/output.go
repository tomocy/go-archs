package output

type UsecaseOutput interface {
	OnError(err error)
}
