package repository

import "github.com/tomocy/archs/domain/model"

type UserRepository interface {
	NextID() model.UserID
	Find(id model.UserID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(user *model.User) error
}
