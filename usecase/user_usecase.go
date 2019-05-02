package usecase

import (
	"github.com/tomocy/archs/usecase/input"
	"github.com/tomocy/archs/usecase/output"
)

type UserUsecase interface {
	RegisterUser(input input.RegisterUserInput, output output.RegisterUserOutput)
}
