package response

import "github.com/tomocy/archs/domain/model"

type AuthenticationUsecaseResponser interface {
	ResponseUser(user *model.User) *UserResponse
}
