package usecase

import (
	"testing"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/db"
	"github.com/tomocy/archs/infra/hash"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name   string
		tester func(t *testing.T)
	}{
		{"success", testRegisterUserSuccessfully},
	}

	for _, test := range tests {
		t.Run(test.name, test.tester)
	}
}

func testRegisterUserSuccessfully(t *testing.T) {
	memory := db.NewMemory()
	bcrypt := hash.NewBcrypt()
	usecase, input, output := prepare(t, memory, bcrypt)

	input.toRegisterUserTester = func() *model.User {
		return &model.User{
			Email:    "aiueo@aiueo.com",
			Password: "aiueo",
		}
	}

	output.expectToBeSuccess()
	output.onUserRegisteredTester = func(t *testing.T, user *model.User) {
		found, err := memory.FindUser(user.ID)
		if err != nil {
			t.Fatalf("failed to find user: %s\n", err)
		}

		assertUser(t, found, user)
	}

	usecase.RegisterUser(input, output)
}

func TestFindUser(t *testing.T) {
	tests := []struct {
		name   string
		tester func(t *testing.T)
	}{
		{"success", testFindUserSuccessfully},
	}

	for _, test := range tests {
		t.Run(test.name, test.tester)
	}
}

func testFindUserSuccessfully(t *testing.T) {
	memory := db.NewMemory()
	bcrypt := hash.NewBcrypt()
	usecase, input, output := prepare(t, memory, bcrypt)

	id := model.UserID("aiueo")
	stored := &model.User{
		ID: id,
	}
	memory.SaveUser(stored)

	input.toFindUserTester = func() model.UserID {
		return id
	}

	output.expectToBeSuccess()
	output.onUserFound = func(t *testing.T, user *model.User) {
		found, err := memory.FindUser(user.ID)
		if err != nil {
			t.Fatalf("failed to find user: %s\n", err)
		}

		assertUser(t, found, user)
	}

	usecase.FindUser(input, output)
}
