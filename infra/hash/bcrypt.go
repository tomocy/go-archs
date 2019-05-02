package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func NewBcrypt() *Bcrypt {
	return new(Bcrypt)
}

type Bcrypt struct {
}

func (b *Bcrypt) Hash(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
