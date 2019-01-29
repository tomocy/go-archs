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
	AuthenticateUser(req *request.AuthenticateUserRequest) (*response.AuthenticateUserResponse, error)
}

type userUsecase struct {
	repository     repository.UserRepository
	responseWriter response.UserResponseWriter
	service        service.UserService
	sessionService service.SessionService
}

func NewUserUsecase(repo repository.UserRepository, w response.UserResponseWriter, s service.UserService, sessService service.SessionService) UserUsecase {
	return &userUsecase{
		repository:     repo,
		responseWriter: w,
		service:        s,
		sessionService: sessService,
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

func (u userUsecase) AuthenticateUser(req *request.AuthenticateUserRequest) (*response.AuthenticateUserResponse, error) {
	user, err := u.repository.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %s", err)
	}
	if err := u.service.ComparePasswords(req.Password, user.Password); err != nil {
		return nil, newIncorrectCredentialError()
	}
	if err := u.sessionService.StoreAuthenticUser(req.ResponseWriter, req.Request, user); err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %s", err)
	}

	return u.responseWriter.WriteAuthenticateUserResponse(user)
}
