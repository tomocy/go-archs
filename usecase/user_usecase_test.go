package usecase

import (
	"testing"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/infra/memory"
	"github.com/tomocy/archs/usecase/request"
)

func TestRegisterUser(t *testing.T) {
	repo := memory.NewUserRepository()
	usecase := NewUserUsecase(
		repo,
		service.NewUserService(repo, new(mockHashService)),
		new(mockHashService),
	)
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
