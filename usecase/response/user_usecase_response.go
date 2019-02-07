package response

import "github.com/tomocy/archs/domain/model"

type UserUsecaseResponser interface {
	ResponseUser(user *model.User) *UserResponse
}

type UserResponse struct {
	ID    string
	Email string
}

func NewUserResponse(id, email string) *UserResponse {
	return &UserResponse{
		ID:    id,
		Email: email,
	}
}
