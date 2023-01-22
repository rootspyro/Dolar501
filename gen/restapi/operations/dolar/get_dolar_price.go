// Code generated by go-swagger; DO NOT EDIT.

package dolar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDolarPriceHandlerFunc turns a function with the right signature into a get dolar price handler
type GetDolarPriceHandlerFunc func(GetDolarPriceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDolarPriceHandlerFunc) Handle(params GetDolarPriceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDolarPriceHandler interface for that can handle valid get dolar price params
type GetDolarPriceHandler interface {
	Handle(GetDolarPriceParams, interface{}) middleware.Responder
}

// NewGetDolarPrice creates a new http.Handler for the get dolar price operation
func NewGetDolarPrice(ctx *middleware.Context, handler GetDolarPriceHandler) *GetDolarPrice {
	return &GetDolarPrice{Context: ctx, Handler: handler}
}

/*
	GetDolarPrice swagger:route GET /dolar/{plataforma} Dolar getDolarPrice

Retorna el precio del dolar en la plataforma indicada
*/
type GetDolarPrice struct {
	Context *middleware.Context
	Handler GetDolarPriceHandler
}

func (o *GetDolarPrice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDolarPriceParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
