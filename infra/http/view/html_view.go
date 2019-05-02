package view

import (
	"net/http"
	"path/filepath"

	"github.com/tomocy/caster"
)

func NewHTML() *HTMLView {
	v := new(HTMLView)
	v.mustParseTemplates()
	return v
}

type HTMLView struct {
	caster caster.Caster
}

func (v *HTMLView) Show(w http.ResponseWriter, name string, data interface{}) error {
	return v.caster.Cast(w, name, data)
}

func (v *HTMLView) mustParseTemplates() {
	var err error
	v.caster, err = caster.New(
		&caster.TemplateSet{
			Filenames: []string{htmlTemplate("master.html")},
		},
	)
	if err != nil {
		panic(err)
	}

	if err := v.caster.ExtendAll(
		map[string]*caster.TemplateSet{},
	); err != nil {
		panic(err)
	}
}

func htmlTemplate(fname string) string {
	projPath, _ := filepath.Abs(".")
	return filepath.Join(projPath, "resource/template/html", fname)
}
