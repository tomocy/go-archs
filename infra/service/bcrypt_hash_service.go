package service

import (
	"github.com/tomocy/archs/domain/service"
	"golang.org/x/crypto/bcrypt"
)

func NewHashService() service.HashService {
	return newBcryptHashService()
}

type bcryptHashService struct {
}

func newBcryptHashService() *bcryptHashService {
	return new(bcryptHashService)
}

func (s bcryptHashService) GenerateHashFromPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hash), err
}

func (s bcryptHashService) ComparePasswords(plain, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
