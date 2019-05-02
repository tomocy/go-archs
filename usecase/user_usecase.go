package usecase

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	uerr "github.com/tomocy/archs/usecase/error"
	"github.com/tomocy/archs/usecase/input"
	"github.com/tomocy/archs/usecase/output"
)

type UserUsecase interface {
	RegisterUser(input input.RegisterUserInput, output output.RegisterUserOutput)
	FindUser(input input.FindUserInput, output output.FindUserOutput)
}

func newUserUsecase(
	repo repository.UserRepository,
	hashServ service.HashService,
) *userUsecase {
	return &userUsecase{
		repo:     repo,
		hashServ: hashServ,
	}
}

type userUsecase struct {
	repo     repository.UserRepository
	hashServ service.HashService
}

func (u *userUsecase) registerUser(input input.RegisterUserInput, output output.RegisterUserOutput) {
	user := input.ToRegisterUser()
	if err := user.AllocateID(u.repo.NextUserID()); err != nil {
		output.OnError(
			wrapValidationError(uerr.KindUserRegistration, "register user", err),
		)
		return
	}
	if err := user.HashPassword(u.hashServ); err != nil {
		// TODO: switch on error
		output.OnError(
			wrapValidationError(uerr.KindUserRegistration, "register user", err),
		)
		return
	}
	if err := user.ValidateSelf(); err != nil {
		output.OnError(
			wrapValidationError(uerr.KindUserRegistration, "register user", err),
		)
		return
	}

	if err := u.repo.SaveUser(user); err != nil {
		// TODO: switch on error
		output.OnError(
			wrapValidationError(uerr.KindUserRegistration, "register user", err),
		)
		return
	}

	output.OnUserRegistered(user)
}

func (u *userUsecase) findUser(input input.FindUserInput, output output.FindUserOutput) {
	id := input.ToFindUser()
	user, err := u.repo.FindUser(id)
	if err != nil {
		// TODO: switch on error
		output.OnError(
			wrapValidationError(uerr.KindUserRegistration, "find user", err),
		)
		return
	}

	output.OnUserFound(user)
}

func wrapValidationError(kind uerr.Kind, did string, err error) error {
	return errors.Wrap(
		uerr.NewValidationError(kind, err.Error()),
		fmt.Sprintf("failed to %s", did),
	)
}
