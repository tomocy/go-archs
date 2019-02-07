package presenter

import "github.com/tomocy/archs/usecase/response"

func NewAuthenticationUsecaseResponser(userResponser response.UserUsecaseResponser) response.AuthenticationUsecaseResponser {
	return &authenticationPresenter{
		UserUsecaseResponser: userResponser,
	}
}

type authenticationPresenter struct {
	response.UserUsecaseResponser
}
