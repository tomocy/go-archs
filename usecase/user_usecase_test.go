package usecase

import (
	"testing"

	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/infra/memory"
	"github.com/tomocy/archs/usecase/request"
)

func TestRegisterUser(t *testing.T) {
	repo := memory.NewUserRepository()
	usecase := NewUserUsecase(repo, presenter.NewUserResponseWriter(), service.NewUserService(repo, new(mockHashService)), new(mockSessionService))
	email := "test@test.com"
	password := "plain"
	test := struct {
		cmd      *request.RegisterUserRequest
		expected *model.User
	}{
		cmd:      request.NewRegisterUserRequest(email, password),
		expected: model.NewUser(email, mockHash),
	}
	tests := []struct {
		name   string
		tester func(t *testing.T)
	}{
		{
			"normal",
			func(t *testing.T) {
				_, err := usecase.RegisterUser(test.cmd)
				if err != nil {
					t.Fatalf("unexpected error: %s\n", err)
				}
			},
		},
		{
			"duplicated email",
			func(t *testing.T) {
				repo.Save(model.NewUser(email, password))
				_, err := usecase.RegisterUser(test.cmd)
				if !IsDuplicatedEmailError(err) {
					t.Errorf("unexpected error: got %s, but expected DuplicatedEmailError", err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tester(t)
		})
	}
}

func TestAuthenticateUser(t *testing.T) {
	repo := memory.NewUserRepository()
	usecase := NewUserUsecase(repo, presenter.NewUserResponseWriter(), service.NewUserService(repo, new(mockHashService)), new(mockSessionService))
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
