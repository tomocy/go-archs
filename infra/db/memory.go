package db

import (
	"errors"

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

	return nil, errors.New("no such user")
}

func (m *Memory) SaveUser(user *model.User) error {
	m.users[user.ID] = user
	return nil
}
