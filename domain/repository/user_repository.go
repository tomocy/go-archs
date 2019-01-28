package repository

import "github.com/tomocy/archs/domain/model"

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	Save(user *model.User) error
}
