// Code generated by go-swagger; DO NOT EDIT.

package dolar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDolarPlatformsHandlerFunc turns a function with the right signature into a get dolar platforms handler
type GetDolarPlatformsHandlerFunc func(GetDolarPlatformsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDolarPlatformsHandlerFunc) Handle(params GetDolarPlatformsParams) middleware.Responder {
	return fn(params)
}

// GetDolarPlatformsHandler interface for that can handle valid get dolar platforms params
type GetDolarPlatformsHandler interface {
	Handle(GetDolarPlatformsParams) middleware.Responder
}

// NewGetDolarPlatforms creates a new http.Handler for the get dolar platforms operation
func NewGetDolarPlatforms(ctx *middleware.Context, handler GetDolarPlatformsHandler) *GetDolarPlatforms {
	return &GetDolarPlatforms{Context: ctx, Handler: handler}
}

/*
	GetDolarPlatforms swagger:route GET /dolar Dolar getDolarPlatforms

Retorna la lista de plataformas del precio del dolar
*/
type GetDolarPlatforms struct {
	Context *middleware.Context
	Handler GetDolarPlatformsHandler
}

func (o *GetDolarPlatforms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDolarPlatformsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}