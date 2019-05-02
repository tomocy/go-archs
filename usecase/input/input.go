package input

import "github.com/tomocy/archs/domain/model"

type RegisterUserInput interface {
	ToRegisterUser() *model.User
}

type FindUserInput interface {
	ToFindUser() model.UserID
}
