package gorilla

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/domain/service"
)

var SessionService service.SessionService = newSessionService()

const sessionKey = "IHAVEAPEN"

type sessionService struct {
	store sessions.Store
}

func newSessionService() *sessionService {
	service := &sessionService{
		store: sessions.NewCookieStore([]byte(sessionKey)),
	}
	service.registerCustomTypes()
	return service
}

func (s sessionService) registerCustomTypes() {
	gob.Register(model.UserID(""))
}

func (s sessionService) StoreAuthenticUser(w http.ResponseWriter, r *http.Request, user *model.User) error {
	sess, err := s.store.Get(r, sessionKey)
	if err != nil {
		return err
	}

	sess.Values["authenticated"] = true
	sess.Values["user_id"] = user.ID
	return sess.Save(r, w)
}
