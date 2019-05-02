package usecase

import (
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/usecase/input"
	"github.com/tomocy/archs/usecase/output"
)

func New(
	userRepo repository.UserRepository,
	hashServ service.HashService,
) *Usecase {
	return &Usecase{
		userUsecase: newUserUsecase(userRepo, hashServ),
	}
}

type Usecase struct {
	userUsecase *userUsecase
}

func (u *Usecase) RegisterUser(input input.RegisterUserInput, output output.RegisterUserOutput) {
	u.userUsecase.registerUser(input, output)
}

func (u *Usecase) FindUser(input input.FindUserInput, output output.FindUserOutput) {
	u.userUsecase.findUser(input, output)
}
