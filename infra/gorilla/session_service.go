package gorilla

import (
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
	return &sessionService{
		store: sessions.NewCookieStore([]byte(sessionKey)),
	}
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
