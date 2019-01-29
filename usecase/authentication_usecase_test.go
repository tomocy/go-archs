package usecase

import (
	"testing"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/memory"
	"github.com/tomocy/archs/usecase/request"
)

func TestAuthenticateUser(t *testing.T) {
	repo := memory.NewUserRepository()
	usecase := NewAuthenticationUsecase(
		repo,
		new(mockHashService),
		new(mockSessionService),
	)
	email := "test@test.com"
	repo.Save(model.NewUser(email, mockHash))
	test := struct {
		cmd      *request.AuthenticateUserRequest
		expected *model.User
	}{
		cmd:      request.NewAuthenticateUserRequest(nil, nil, email, mockPlain),
		expected: model.NewUser(email, mockHash),
	}
	tests := []struct {
		name   string
		tester func(t *testing.T)
	}{
		{
			"normal",
			func(t *testing.T) {
				_, err := usecase.AuthenticateUser(test.cmd)
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
			},
		},
		{
			"incorrent password",
			func(t *testing.T) {
				test.cmd.Password = ""
				_, err := usecase.AuthenticateUser(test.cmd)
				if !IsIncorrectCredentialError(err) {
					t.Errorf("unexpected error: got %v, but expected IncorrectCredentialError\n", err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.tester)
	}
}
