package output

import "github.com/tomocy/archs/domain/model"

type UsecaseOutput interface {
	OnError(err error)
}

type RegisterUserOutput interface {
	UsecaseOutput
	OnUserRegistered(user *model.User)
}

type FindUserOutput interface {
	UsecaseOutput
	OnUserFound(user *model.User)
}
