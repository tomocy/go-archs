package usecase

import (
	"testing"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/db"
	"github.com/tomocy/archs/infra/hash"
)

type testInput struct {
	toRegisterUserTester func() *model.User
	toFindUserTester     func() model.UserID
}

func (i *testInput) ToRegisterUser() *model.User {
	return i.toRegisterUserTester()
}

func (i *testInput) ToFindUser() model.UserID {
	return i.toFindUserTester()
}

type testOutput struct {
	t                              *testing.T
	onUserRegisteredTester         func(t *testing.T, user *model.User)
	onUserRegistrationFailedTester func(t *testing.T, err error)
	onUserFound                    func(t *testing.T, user *model.User)
	onUserFindingFailedTester      func(t *testing.T, err error)
}

func (o *testOutput) OnUserRegistrationFailed(err error) {
	o.onUserRegistrationFailedTester(o.t, err)
}

func (o *testOutput) OnUserRegistered(user *model.User) {
	o.onUserRegisteredTester(o.t, user)
}

func (o *testOutput) OnUserFindingFailed(err error) {
	o.onUserFindingFailedTester(o.t, err)
}

func (o *testOutput) OnUserFound(user *model.User) {
	o.onUserFound(o.t, user)
}

func prepare(
	t *testing.T,
	memory *db.Memory,
	bcrypt *hash.Bcrypt,
) (*Usecase, *testInput, *testOutput) {
	return New(memory, bcrypt), new(testInput), &testOutput{t: t}
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
