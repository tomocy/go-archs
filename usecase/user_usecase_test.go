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
	usecase := New(memory, bcrypt)
	input := new(testInput)
	output := &testOutput{t: t}
	input.toRegisterUserTester = func() *model.User {
		return &model.User{
			Email:    "aiueo@aiueo.com",
			Password: "aiueo",
		}
	}
	output.onErrorTester = func(t *testing.T, err error) {
		t.Fatalf("onError was called despite the fact that this test is expected to be success: %s\n", err)
	}
	output.onUserRegisteredTester = func(t *testing.T, user *model.User) {
		found, err := memory.FindUser(user.ID)
		if err != nil {
			t.Fatalf("failed to find user: %s\n", err)
		}
		assertUser(t, found, user)
	}

	usecase.RegisterUser(input, output)
}
