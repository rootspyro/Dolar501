// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DolarPrice dolar price
//
// swagger:model DolarPrice
type DolarPrice struct {

	// plataforma
	Plataforma string `json:"plataforma,omitempty"`

	// precio ves
	PrecioVes float64 `json:"precio_ves,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this dolar price
func (m *DolarPrice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dolar price based on context it is used
func (m *DolarPrice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DolarPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DolarPrice) UnmarshalBinary(b []byte) error {
	var res DolarPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
