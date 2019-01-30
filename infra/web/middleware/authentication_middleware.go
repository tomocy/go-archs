package middleware

import (
	"net/http"

	"github.com/tomocy/archs/infra/gorilla"
)

func Deauthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if gorilla.SessionService.HasAuthenticUser(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
