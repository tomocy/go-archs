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
	s.store["user_id"] = string(user.ID)

	return nil
}

func (s *mockSessionService) HasAuthenticUser(r *http.Request) bool {
	if s.store == nil {
		s.store = make(map[string]interface{})
	}

	authenticated, ok := s.store["authenticated"].(bool)
	return authenticated && ok
}

func (s *mockSessionService) GetAuthenticUserID(r *http.Request) string {
	if s.store == nil {
		s.store = make(map[string]interface{})
	}

	userID, ok := s.store["user_id"].(string)
	if !ok {
		return ""
	}

	return userID
}
