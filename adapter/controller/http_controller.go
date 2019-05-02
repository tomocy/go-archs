package controller

import (
	"net/http"
)

type HTTPController struct {
	request *http.Request
}

func NewHTTPController(r *http.Request) *HTTPController {
	return &HTTPController{
		request: r,
	}
}
