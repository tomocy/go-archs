package service

import (
	"net/http"

	"github.com/tomocy/archs/domain/model"
)

type SessionService interface {
	StoreAuthenticUser(w http.ResponseWriter, r *http.Request, user *model.User) error
	HasAuthenticUser(r *http.Request) bool
	GetAuthenticUserID(r *http.Request) string
}
