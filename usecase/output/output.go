package output

import "github.com/tomocy/archs/domain/model"

type UsecaseOutput interface {
	OnError(err error)
}

type RegisterUserOutput interface {
	OnUserRegistered(user *model.User)
	OnUserRegistrationFailed(err error)
}

type FindUserOutput interface {
	UsecaseOutput
	OnUserFound(user *model.User)
}
