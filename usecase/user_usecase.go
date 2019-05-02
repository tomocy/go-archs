package usecase

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
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
		output.OnError(wrapError(err, "register user"))
		return
	}
	if err := user.HashPassword(u.hashServ); err != nil {
		output.OnError(wrapError(err, "register user"))
		return
	}
	if err := user.ValidateSelf(); err != nil {
		output.OnError(wrapError(err, "register user"))
		return
	}

	if err := u.repo.SaveUser(user); err != nil {
		output.OnError(wrapError(err, "register user"))
		return
	}

	output.OnUserRegistered(user)
}

func (u *userUsecase) findUser(input input.FindUserInput, output output.FindUserOutput) {
	id := input.ToFindUser()
	user, err := u.repo.FindUser(id)
	if err != nil {
		output.OnError(wrapError(err, "find user"))
		return
	}

	output.OnUserFound(user)
}

func wrapError(err error, did string) error {
	return errors.Wrap(err, fmt.Sprintf("failed to %s", did))
}
