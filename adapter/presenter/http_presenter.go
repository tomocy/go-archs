package presenter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/http/route"
	uerr "github.com/tomocy/archs/usecase/error"
)

func NewHTTPPresenter(w http.ResponseWriter, r *http.Request) *HTTPPresenter {
	return &HTTPPresenter{
		respWriter: w,
		request:    r,
	}
}

type HTTPPresenter struct {
	respWriter http.ResponseWriter
	request    *http.Request
}

func (p *HTTPPresenter) OnUserRegistered(user *model.User) {
	dest := fmt.Sprintf("%s/%s", route.Web.Route("user.show"), user.ID)
	p.redirect(dest)
}

func (p *HTTPPresenter) OnError(err error) {
	switch {
	case uerr.InUserRegistration(err):
		p.onUserRegistrationError(err)
	default:
		p.logInternalServerError("unknown", err)
	}
}

func (p *HTTPPresenter) onUserRegistrationError(err error) {
	switch {
	case uerr.InInput(err):
		// TODO: redirect to proper location with error message
	default:
		p.logInternalServerError("user registration", err)
	}
}

func (p *HTTPPresenter) redirect(dest string) {
	http.Redirect(p.respWriter, p.request, dest, http.StatusSeeOther)
}

func (p *HTTPPresenter) logInternalServerError(in string, err error) {
	log.Printf("failed to deal with unknown error in %s: %v\n", in, err)
	p.respWriter.WriteHeader(http.StatusInternalServerError)
}
