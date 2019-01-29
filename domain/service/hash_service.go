package service

type HashService interface {
	GenerateHashFromPassword(plan string) (string, error)
	ComparePasswords(plain, hash string) error
}

type hashService struct {
}
