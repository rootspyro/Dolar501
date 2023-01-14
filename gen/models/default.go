// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Default default
//
// swagger:model default
type Default struct {

	// data
	Data string `json:"data,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this default
func (m *Default) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this default based on context it is used
func (m *Default) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Default) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Default) UnmarshalBinary(b []byte) error {
	var res Default
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}