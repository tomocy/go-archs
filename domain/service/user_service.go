package service

import (
	"fmt"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
)

type UserService interface {
	RegisterUser(id model.UserID, email, password string) (*model.User, error)
}

type userService struct {
	repository  repository.UserRepository
	hashService HashService
}

func NewUserService(repo repository.UserRepository, hashService HashService) UserService {
	return &userService{
		repository:  repo,
		hashService: hashService,
	}
}

func (s userService) RegisterUser(id model.UserID, email, password string) (*model.User, error) {
	if err := s.checkIfEmailIsUnique(email); err != nil {
		return nil, err
	}

	hash, err := s.hashService.GenerateHashFromPassword(password)
	if err != nil {
		return nil, err
	}

	return model.NewUser(id, email, hash), nil
}

func (s userService) checkIfEmailIsUnique(email string) error {
	user, _ := s.repository.FindByEmail(email)
	if user != nil {
		return fmt.Errorf("email is not unique: %s", email)
	}

	return nil
}
