package service

type HashService interface {
	Hash(plain string) (string, error)
}
