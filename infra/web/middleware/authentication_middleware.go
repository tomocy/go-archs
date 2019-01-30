package middleware

import (
	"net/http"

	"github.com/tomocy/archs/infra/session"
)

func Deauthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session.SessionService.HasAuthenticUser(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
