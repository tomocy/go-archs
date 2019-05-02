package session

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/tomocy/sensei"
)

func SetErrorMessage(w http.ResponseWriter, r *http.Request, msg string) {
	if err := manager.SetFlash(w, r, flashError, msg); err != nil {
		logError("set error", err)
	}
}

func GetErrorMessage(w http.ResponseWriter, r *http.Request) string {
	flashes, err := manager.GetFlashes(w, r, flashError)
	if err != nil {
		logError("get error message", err)
		return ""
	}

	for _, flash := range flashes {
		if s, ok := flash.(string); ok {
			return s
		}
	}

	return ""
}

func IsSessionValid(r *http.Request) bool {
	_, err := manager.Session(r)
	return err == nil
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	if err := manager.Delete(w, r); err != nil {
		logError("delete session", err)
	}
}

func logError(did string, err error) {
	log.Printf("failed to %s: %s\n", did, err)
}

var manager = sensei.New(store, sessionKey)

var store = sessions.NewCookieStore(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

const sessionKey = "IHaveAPen"

const (
	flashError = "errors"
)
