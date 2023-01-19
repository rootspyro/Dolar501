// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rootspyro/Dolar501/gen/models"
)

// GetAuthTokenOKCode is the HTTP code returned for type GetAuthTokenOK
const GetAuthTokenOKCode int = 200

/*
GetAuthTokenOK OK

swagger:response getAuthTokenOK
*/
type GetAuthTokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccessResponse `json:"body,omitempty"`
}

// NewGetAuthTokenOK creates GetAuthTokenOK with default headers values
func NewGetAuthTokenOK() *GetAuthTokenOK {

	return &GetAuthTokenOK{}
}

// WithPayload adds the payload to the get auth token o k response
func (o *GetAuthTokenOK) WithPayload(payload *models.AccessResponse) *GetAuthTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth token o k response
func (o *GetAuthTokenOK) SetPayload(payload *models.AccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetAuthTokenDefault respuesta generica

swagger:response getAuthTokenDefault
*/
type GetAuthTokenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Default `json:"body,omitempty"`
}

// NewGetAuthTokenDefault creates GetAuthTokenDefault with default headers values
func NewGetAuthTokenDefault(code int) *GetAuthTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAuthTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get auth token default response
func (o *GetAuthTokenDefault) WithStatusCode(code int) *GetAuthTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get auth token default response
func (o *GetAuthTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get auth token default response
func (o *GetAuthTokenDefault) WithPayload(payload *models.Default) *GetAuthTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get auth token default response
func (o *GetAuthTokenDefault) SetPayload(payload *models.Default) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}