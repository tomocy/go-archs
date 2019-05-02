package usecase

import (
	"testing"

	"github.com/tomocy/archs/domain/model"
)

type testInput struct {
	toRegisterUserTester func() *model.User
}

func (i *testInput) ToRegisterUser() *model.User {
	return i.toRegisterUserTester()
}

type testOutput struct {
	t                      *testing.T
	onErrorTester          func(t *testing.T, err error)
	onUserRegisteredTester func(t *testing.T, user *model.User)
}

func (o *testOutput) OnError(err error) {
	o.onErrorTester(o.t, err)
}

func (o *testOutput) OnUserRegistered(user *model.User) {
	o.onUserRegisteredTester(o.t, user)
}

func assertUser(t *testing.T, actual, expected *model.User) {
	if actual.ID != expected.ID {
		errorf(t, actual.ID, expected.ID)
	}
	if actual.Email != expected.Email {
		errorf(t, actual.Email, expected.Email)
	}
	if actual.Password != expected.Password {
		errorf(t, actual.Password, expected.Password)
	}
}

func errorf(t *testing.T, actual, expected interface{}) {
	t.Errorf("got %v, expected %v\n", actual, expected)
}
