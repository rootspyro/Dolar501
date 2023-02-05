// Code generated by go-swagger; DO NOT EDIT.

package dolar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rootspyro/Dolar501/gen/models"
)

// GetCurrencyPlatformsOKCode is the HTTP code returned for type GetCurrencyPlatformsOK
const GetCurrencyPlatformsOKCode int = 200

/*
GetCurrencyPlatformsOK OK

swagger:response getCurrencyPlatformsOK
*/
type GetCurrencyPlatformsOK struct {

	/*
	  In: Body
	*/
	Payload *models.DolarPlatforms `json:"body,omitempty"`
}

// NewGetCurrencyPlatformsOK creates GetCurrencyPlatformsOK with default headers values
func NewGetCurrencyPlatformsOK() *GetCurrencyPlatformsOK {

	return &GetCurrencyPlatformsOK{}
}

// WithPayload adds the payload to the get currency platforms o k response
func (o *GetCurrencyPlatformsOK) WithPayload(payload *models.DolarPlatforms) *GetCurrencyPlatformsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get currency platforms o k response
func (o *GetCurrencyPlatformsOK) SetPayload(payload *models.DolarPlatforms) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrencyPlatformsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetCurrencyPlatformsDefault respuesta generica

swagger:response getCurrencyPlatformsDefault
*/
type GetCurrencyPlatformsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Default `json:"body,omitempty"`
}

// NewGetCurrencyPlatformsDefault creates GetCurrencyPlatformsDefault with default headers values
func NewGetCurrencyPlatformsDefault(code int) *GetCurrencyPlatformsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCurrencyPlatformsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get currency platforms default response
func (o *GetCurrencyPlatformsDefault) WithStatusCode(code int) *GetCurrencyPlatformsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get currency platforms default response
func (o *GetCurrencyPlatformsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get currency platforms default response
func (o *GetCurrencyPlatformsDefault) WithPayload(payload *models.Default) *GetCurrencyPlatformsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get currency platforms default response
func (o *GetCurrencyPlatformsDefault) SetPayload(payload *models.Default) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrencyPlatformsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
