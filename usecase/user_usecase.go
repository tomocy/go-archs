package usecase

import (
	"fmt"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/usecase/request"
)

type UserUsecase interface {
	RegisterUser(req *request.RegisterUserRequest) (*model.User, error)
}

type userUsecase struct {
	repository  repository.UserRepository
	userService service.UserService
	hashService service.HashService
}

func NewUserUsecase(
	repo repository.UserRepository,
	userService service.UserService,
	hashService service.HashService,
) UserUsecase {
	return &userUsecase{
		repository:  repo,
		userService: userService,
		hashService: hashService,
	}
}

func (u userUsecase) RegisterUser(req *request.RegisterUserRequest) (*model.User, error) {
	user, err := u.userService.RegisterUser(u.repository.NextID(), req.Email, req.Password)
	if err != nil {
		return nil, newDuplicatedEmailError(req.Email)
	}
	if err := u.repository.Save(user); err != nil {
		return nil, fmt.Errorf("failed to register user: %s", err)
	}

	return user, nil
}
