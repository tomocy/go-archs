package presenter

import (
	"net/http"
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
