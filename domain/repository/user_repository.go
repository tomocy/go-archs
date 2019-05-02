package repository

import "github.com/tomocy/archs/domain/model"

type UserRepository interface {
	NextUserID() model.UserID
	SaveUser(user *model.User) error
}
