package request

type RegisterUserRequest struct {
	Email    string
	Password string
}

func NewRegisterUserRequest(email, password string) *RegisterUserRequest {
	return &RegisterUserRequest{
		Email:    email,
		Password: password,
	}
}
