package presenter

import "github.com/tomocy/archs/domain/model"

type UserPresenter interface {
	PresentUser(user *model.User) *UserPresent
}

func NewUserPresenter() UserPresenter {
	return new(userPresenter)
}

type userPresenter struct {
}

type UserPresent struct {
	ID    string
	Email string
}

func (p userPresenter) PresentUser(user *model.User) *UserPresent {
	return &UserPresent{
		ID:    string(user.ID),
		Email: user.Email,
	}
}
