package usecase

import (
	"fmt"

	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/usecase/request"
	"github.com/tomocy/archs/usecase/response"
)

type UserUsecase interface {
	RegisterUser(req *request.RegisterUserRequest) (*response.UserResponse, error)
}

type userUsecase struct {
	responser   response.UserUsecaseResponser
	repository  repository.UserRepository
	userService service.UserService
	hashService service.HashService
}

func NewUserUsecase(
	responser response.UserUsecaseResponser,
	repo repository.UserRepository,
	userService service.UserService,
	hashService service.HashService,
) UserUsecase {
	return &userUsecase{
		responser:   responser,
		repository:  repo,
		userService: userService,
		hashService: hashService,
	}
}

func (u userUsecase) RegisterUser(req *request.RegisterUserRequest) (*response.UserResponse, error) {
	user, err := u.userService.RegisterUser(req.Email, req.Password)
	if err != nil {
		return nil, newDuplicatedEmailError(req.Email)
	}
	if err := u.repository.Save(user); err != nil {
		return nil, fmt.Errorf("failed to register user: %s", err)
	}

	return u.responser.ResponseUser(user)
}
