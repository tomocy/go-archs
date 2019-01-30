package session

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/service"
)

var SessionService service.SessionService = newGorillaSessionService()

const sessionKey = "IHAVEAPEN"

type gorillaSessionService struct {
	store sessions.Store
}

func newGorillaSessionService() *gorillaSessionService {
	service := &gorillaSessionService{
		store: sessions.NewCookieStore([]byte(sessionKey)),
	}
	service.registerCustomTypes()
	return service
}

func (s gorillaSessionService) registerCustomTypes() {
	gob.Register(model.UserID(""))
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

func (s gorillaSessionService) HasAuthenticUser(r *http.Request) bool {
	sess, err := s.store.Get(r, sessionKey)
	if err != nil {
		return false
	}

	authenticated, ok := sess.Values["authenticated"].(bool)
	return authenticated && ok
}
