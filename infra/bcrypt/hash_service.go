package bcrypt

import (
	"github.com/tomocy/archs/domain/service"
	"golang.org/x/crypto/bcrypt"
)

var HashService service.HashService = new(hashService)

type hashService struct {
}

func (s hashService) GenerateHashFromPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hash), err
}

func (s hashService) ComparePasswords(plain, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
