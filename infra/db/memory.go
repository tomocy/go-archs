package db

import (
	derr "github.com/tomocy/archs/domain/error"
	"github.com/tomocy/archs/domain/model"
)

func NewMemory() *Memory {
	return &Memory{
		users: make(map[model.UserID]*model.User),
	}
}

type Memory struct {
	users map[model.UserID]*model.User
}

func (m *Memory) NextUserID() model.UserID {
	return model.UserID(generateULID())
}

func (m *Memory) FindUser(id model.UserID) (*model.User, error) {
	if stored, ok := m.users[id]; ok {
		return stored, nil
	}

	return nil, derr.NewValidationError("no such user")
}

func (m *Memory) SaveUser(user *model.User) error {
	if m.hasSameEmail(user.Email) {
		return derr.NewValidationError("duplicated email")
	}

	m.users[user.ID] = user

	return nil
}

func (m *Memory) hasSameEmail(email string) bool {
	for _, stored := range m.users {
		if stored.Email == email {
			return true
		}
	}

	return false
}
