package usecase

import (
	"fmt"

	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/usecase/request"
	"github.com/tomocy/archs/usecase/response"
)

type UserUsecase interface {
	RegisterUser(req *request.RegisterUserRequest) (*response.RegisterUserResponse, error)
}

type userUsecase struct {
	repository     repository.UserRepository
	responseWriter response.UserResponseWriter
	service        service.UserService
}

func NewUserUsecase(repo repository.UserRepository, w response.UserResponseWriter, s service.UserService) UserUsecase {
	return &userUsecase{
		repository:     repo,
		responseWriter: w,
		service:        s,
	}
}

func (u userUsecase) RegisterUser(req *request.RegisterUserRequest) (*response.RegisterUserResponse, error) {
	user, err := u.service.RegisterUser(req.Email, req.Password)
	if err != nil {
		return nil, newDuplicatedEmailError(req.Email)
	}
	if err := u.repository.Save(user); err != nil {
		return nil, fmt.Errorf("failed to register user: %s", err)
	}

	return u.responseWriter.WriteRegisterUserResponse(user)
}
