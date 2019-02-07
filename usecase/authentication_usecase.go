package usecase

import (
	"fmt"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/usecase/request"
	"github.com/tomocy/archs/usecase/response"
)

type AuthenticationUsecase interface {
	AuthenticateUser(req *request.AuthenticateUserRequest) (*response.UserResponse, error)
	GetAuthenticUserID(req *request.GetAuthenticUserIDRequest) model.UserID
}

type authenticationUsecase struct {
	responser      response.AuthenticationUsecaseResponser
	userRepository repository.UserRepository
	hashService    service.HashService
	sessionService service.SessionService
}

func NewAuthenticationUsecase(
	responser response.AuthenticationUsecaseResponser,
	userRepo repository.UserRepository,
	hashService service.HashService,
	sessService service.SessionService,
) AuthenticationUsecase {
	return &authenticationUsecase{
		responser:      responser,
		userRepository: userRepo,
		hashService:    hashService,
		sessionService: sessService,
	}
}

func (u authenticationUsecase) AuthenticateUser(req *request.AuthenticateUserRequest) (*response.UserResponse, error) {
	user, err := u.userRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %s", err)
	}
	if err := u.hashService.ComparePasswords(req.Password, user.Password); err != nil {
		return nil, newIncorrectCredentialError()
	}
	if err := u.sessionService.StoreAuthenticUser(req.ResponseWriter, req.Request, user); err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %s", err)
	}

	return u.responser.ResponseUser(user), nil
}

func (u authenticationUsecase) GetAuthenticUserID(req *request.GetAuthenticUserIDRequest) model.UserID {
	return model.UserID(u.sessionService.GetAuthenticUserID(req.Request))
}
