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

func (p *Presenter) ShowUserRegistrationForm() {
	if err := p.view.Show(p.respWriter, "user.new", nil); err != nil {
		p.logInternalServerError("show user registration form", err)
	}
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

func (p *Presenter) logInternalServerError(did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	p.respWriter.WriteHeader(http.StatusInternalServerError)
}

func (p *Presenter) logUnknownError(did string, err error) {
	log.Printf("failed to deal with unknown error in %s: %v\n", did, err)
	p.respWriter.WriteHeader(http.StatusInternalServerError)
}
