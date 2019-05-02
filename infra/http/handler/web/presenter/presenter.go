package presenter

import (
	"fmt"
	"log"
	"net/http"

	derr "github.com/tomocy/archs/domain/error"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/http/route"
	"github.com/tomocy/archs/infra/http/view"
)

func New(view view.View, w http.ResponseWriter, r *http.Request) *Presenter {
	return &Presenter{
		view:       view,
		respWriter: w,
		request:    r,
	}
}

type Presenter struct {
	view       view.View
	respWriter http.ResponseWriter
	request    *http.Request
}

func (p *Presenter) OnUserRegistered(user *model.User) {
	dest := fmt.Sprintf("%s/%s", route.Web.Route("user.show"), user.ID)
	log.Printf("register user successfully: %v\n", user)
	p.redirect(dest)
}

func (p *Presenter) OnUserRegistrationFailed(err error) {
	switch {
	case derr.InInput(err):
		p.redirect(route.Web.Route("user.new").String())
	default:
		p.logUnknownError("user registration", err)
	}
}

func (p *Presenter) OnUserFindingFailed(err error) {
	switch {
	case derr.InInput(err):
		p.respWriter.WriteHeader(http.StatusNotFound)
	default:
		p.logUnknownError("user finding", err)
	}
}

func (p *Presenter) redirect(dest string) {
	http.Redirect(p.respWriter, p.request, dest, http.StatusSeeOther)
}

func (p *Presenter) logUnknownError(did string, err error) {
	p.logInternalServerError("failed to deal with unknown error in %s: %v\n", did, err)
}

func (p *Presenter) logInternalServerError(format string, a ...interface{}) {
	log.Printf(format, a...)
	p.respWriter.WriteHeader(http.StatusInternalServerError)
}

func logInternalServerError(w http.ResponseWriter, did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	w.WriteHeader(http.StatusInternalServerError)
}
