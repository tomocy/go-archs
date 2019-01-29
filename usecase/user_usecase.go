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
	AuthenticateUser(req *request.AuthenticateUserRequest) (*model.User, error)
}

type userUsecase struct {
	repository     repository.UserRepository
	userService    service.UserService
	hashService    service.HashService
	sessionService service.SessionService
}

func NewUserUsecase(
	repo repository.UserRepository,
	userService service.UserService,
	hashService service.HashService,
	sessService service.SessionService,
) UserUsecase {
	return &userUsecase{
		repository:     repo,
		userService:    userService,
		hashService:    hashService,
		sessionService: sessService,
	}
}

func (u userUsecase) RegisterUser(req *request.RegisterUserRequest) (*model.User, error) {
	user, err := u.userService.RegisterUser(req.Email, req.Password)
	if err != nil {
		return nil, newDuplicatedEmailError(req.Email)
	}
	if err := u.repository.Save(user); err != nil {
		return nil, fmt.Errorf("failed to register user: %s", err)
	}

	return user, nil
}

func (u userUsecase) AuthenticateUser(req *request.AuthenticateUserRequest) (*model.User, error) {
	user, err := u.repository.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %s", err)
	}
	if err := u.hashService.ComparePasswords(req.Password, user.Password); err != nil {
		return nil, newIncorrectCredentialError()
	}
	if err := u.sessionService.StoreAuthenticUser(req.ResponseWriter, req.Request, user); err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %s", err)
	}

	return user, nil
}
