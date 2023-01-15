// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DolarPlatform dolar platform
//
// swagger:model DolarPlatform
type DolarPlatform struct {

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// plataforma
	Plataforma string `json:"plataforma,omitempty"`
}

// Validate validates this dolar platform
func (m *DolarPlatform) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dolar platform based on context it is used
func (m *DolarPlatform) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DolarPlatform) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DolarPlatform) UnmarshalBinary(b []byte) error {
	var res DolarPlatform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}