// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAuthTokenHandlerFunc turns a function with the right signature into a get auth token handler
type GetAuthTokenHandlerFunc func(GetAuthTokenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAuthTokenHandlerFunc) Handle(params GetAuthTokenParams) middleware.Responder {
	return fn(params)
}

// GetAuthTokenHandler interface for that can handle valid get auth token params
type GetAuthTokenHandler interface {
	Handle(GetAuthTokenParams) middleware.Responder
}

// NewGetAuthToken creates a new http.Handler for the get auth token operation
func NewGetAuthToken(ctx *middleware.Context, handler GetAuthTokenHandler) *GetAuthToken {
	return &GetAuthToken{Context: ctx, Handler: handler}
}

/*
	GetAuthToken swagger:route GET /auth/callback Auth getAuthToken

retorna el token de acceso
*/
type GetAuthToken struct {
	Context *middleware.Context
	Handler GetAuthTokenHandler
}

func (o *GetAuthToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAuthTokenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}