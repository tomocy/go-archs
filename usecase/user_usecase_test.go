package usecase

import (
	"testing"

	"github.com/pkg/errors"
	derr "github.com/tomocy/archs/domain/error"
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
		{"fail bacause of empty email", testRegisterUserWithEmptyEmail},
		{"fail bacause of empty password", testRegisterUserWithEmptyPassword},
		{"fail bacause of duplicated email", testRegisterUserWithDuplicatedEmail},
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

	output.expectUserRegistrationToBeSuccess()
	output.onUserRegisteredTester = func(t *testing.T, user *model.User) {
		found, err := memory.FindUser(user.ID)
		if err != nil {
			t.Fatalf("failed to find user: %s\n", err)
		}

		assertUser(t, found, user)
	}

	usecase.RegisterUser(input, output)
}

func testRegisterUserWithEmptyEmail(t *testing.T) {
	memory := db.NewMemory()
	bcrypt := hash.NewBcrypt()
	usecase, input, output := prepare(t, memory, bcrypt)

	input.toRegisterUserTester = func() *model.User {
		return &model.User{
			Email:    "",
			Password: "aiueo",
		}
	}

	expectInputErrorInUserRegistration(output)

	usecase.RegisterUser(input, output)
}

func testRegisterUserWithEmptyPassword(t *testing.T) {
	memory := db.NewMemory()
	bcrypt := hash.NewBcrypt()
	usecase, input, output := prepare(t, memory, bcrypt)

	input.toRegisterUserTester = func() *model.User {
		return &model.User{
			Email:    "aiueo@aiueo.com",
			Password: "",
		}
	}

	expectInputErrorInUserRegistration(output)

	usecase.RegisterUser(input, output)
}

func expectInputErrorInUserRegistration(output *testOutput) {
	output.onUserRegistrationFailedTester = func(t *testing.T, err error) {
		cause := errors.Cause(err)
		if !derr.InInput(cause) {
			t.Errorf("unexpected error was returned instead of internal error: %T", cause)
		}
	}
	output.onUserRegisteredTester = func(t *testing.T, _ *model.User) {
		t.Fatalf("OnUserRegistered was called despite the fact that this test is not expected to be success")
	}
}

func testRegisterUserWithDuplicatedEmail(t *testing.T) {
	memory := db.NewMemory()
	bcrypt := hash.NewBcrypt()
	usecase, input, output := prepare(t, memory, bcrypt)

	email := "aiueo@aiueo.com"
	stored := &model.User{
		Email: email,
	}
	memory.SaveUser(stored)

	input.toRegisterUserTester = func() *model.User {
		return &model.User{
			Email:    email,
			Password: "aiueo",
		}
	}

	expectInputErrorInUserRegistration(output)

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

	output.expectUserFindingToBeSuccess()
	output.onUserFound = func(t *testing.T, user *model.User) {
		found, err := memory.FindUser(user.ID)
		if err != nil {
			t.Fatalf("failed to find user: %s\n", err)
		}

		assertUser(t, found, user)
	}

	usecase.FindUser(input, output)
}
