package service

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/service"
)

func NewSessionService() service.SessionService {
	return session
}

var session service.SessionService = newGorillaSessionService()

const sessionKey = "IHAVEAPEN"

type gorillaSessionService struct {
	store sessions.Store
}

func newGorillaSessionService() *gorillaSessionService {
	return &gorillaSessionService{
		store: sessions.NewCookieStore([]byte(sessionKey)),
	}
}

func (s gorillaSessionService) StoreAuthenticUser(w http.ResponseWriter, r *http.Request, user *model.User) error {
	sess, err := s.store.Get(r, sessionKey)
	if err != nil {
		return err
	}

	sess.Values["authenticated"] = true
	sess.Values["user_id"] = user.ID
	return sess.Save(r, w)
}
