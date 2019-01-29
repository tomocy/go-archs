package usecase

import (
	"errors"
	"net/http"

	"github.com/tomocy/archs/domain/model"
)

var mockPlain = "plain"
var mockHash = "hash"

type mockHashService struct {
}

func (s mockHashService) GenerateHashFromPassword(plain string) (string, error) {
	return mockHash, nil
}

func (s mockHashService) ComparePasswords(plain, hash string) error {
	if plain == mockPlain && hash == mockHash {
		return nil
	}

	return errors.New("incorrect password")
}

type mockSessionService struct {
	store map[string]interface{}
}

func (s *mockSessionService) StoreAuthenticUser(w http.ResponseWriter, r *http.Request, user *model.User) error {
	if s.store == nil {
		s.store = make(map[string]interface{})
	}

	s.store["authenticated"] = true
	s.store["user_id"] = user.ID

	return nil
}
