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
	log.Printf("register user successfully: %v\n", user)
	p.redirect(dest)
}

func (p *HTTPPresenter) OnUserRegistrationFailed(err error) {
	switch {
	case uerr.InInput(err):
		log.Printf("input error was occured in user registration: %s\n", err)
		// TODO: redirect to proper location with error message
	default:
		p.logInternalServerError("user registration", err)
	}
}

func (p *HTTPPresenter) OnUserFindingFailed(err error) {
	switch {
	case uerr.InInput(err):
		log.Printf("input error was occured in user finding: %s\n", err)
		p.respWriter.WriteHeader(http.StatusNotFound)
	default:
		p.logInternalServerError("user finding", err)
	}
}

func (p *HTTPPresenter) redirect(dest string) {
	http.Redirect(p.respWriter, p.request, dest, http.StatusSeeOther)
}

func (p *HTTPPresenter) logInternalServerError(did string, err error) {
	log.Printf("failed to deal with unknown error in %s: %v\n", did, err)
	p.respWriter.WriteHeader(http.StatusInternalServerError)
}
