package service

type HashService interface {
	GenerateHashFromPassword(plan string) (string, error)
}

type hashService struct {
}
