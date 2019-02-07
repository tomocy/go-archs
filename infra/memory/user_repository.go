package memory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/repository"
)

var UserRepository repository.UserRepository = NewUserRepository()

func NewUserRepository() repository.UserRepository {
	return new(userRepository)
}

type userRepository struct {
	users []*model.User
}

func (r userRepository) NextID() model.UserID {
	return model.UserID(uuid.New().String())
}

func (r userRepository) Find(id model.UserID) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &model.User{
				ID:       user.ID,
				Email:    user.Email,
				Password: user.Password,
			}, nil
		}
	}

	return nil, fmt.Errorf("no such user whose id is %s", id)
}

func (r userRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &model.User{
				ID:       user.ID,
				Email:    user.Email,
				Password: user.Password,
			}, nil
		}
	}

	return nil, fmt.Errorf("no user whose email is %s found", email)
}

func (r *userRepository) Save(user *model.User) error {
	r.users = append(r.users, user)
	return nil
}
