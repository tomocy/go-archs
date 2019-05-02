package controller

import (
	"net/http"

	"github.com/tomocy/archs/domain/model"
)

type HTTPController struct {
	request *http.Request
}

func NewHTTPController(r *http.Request) *HTTPController {
	return &HTTPController{
		request: r,
	}
}

func (c *HTTPController) ToRegisterUser() *model.User {
	email, password := c.request.FormValue("email"), c.request.FormValue("password")
	return &model.User{
		Email:    email,
		Password: password,
	}
}
