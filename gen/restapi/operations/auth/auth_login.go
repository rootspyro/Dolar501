// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AuthLoginHandlerFunc turns a function with the right signature into a auth login handler
type AuthLoginHandlerFunc func(AuthLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthLoginHandlerFunc) Handle(params AuthLoginParams) middleware.Responder {
	return fn(params)
}

// AuthLoginHandler interface for that can handle valid auth login params
type AuthLoginHandler interface {
	Handle(AuthLoginParams) middleware.Responder
}

// NewAuthLogin creates a new http.Handler for the auth login operation
func NewAuthLogin(ctx *middleware.Context, handler AuthLoginHandler) *AuthLogin {
	return &AuthLogin{Context: ctx, Handler: handler}
}

/*
	AuthLogin swagger:route GET /auth/login Auth authLogin

redirige al usuario al inicio de sesion en github
*/
type AuthLogin struct {
	Context *middleware.Context
	Handler AuthLoginHandler
}

func (o *AuthLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAuthLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
