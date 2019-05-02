package repository

import "github.com/tomocy/archs/domain/model"

type UserRepository interface {
	NextUserID() model.UserID
	FindUser(id model.UserID) (*model.User, error)
	SaveUser(user *model.User) error
}
