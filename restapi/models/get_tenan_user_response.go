// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetTenanUserResponse get tenan user response
//
// swagger:model GetTenanUserResponse
type GetTenanUserResponse struct {

	// email
	Email string `json:"Email"`

	// message
	Message string `json:"Message"`

	// name
	Name string `json:"Name"`

	// status code
	StatusCode int64 `json:"StatusCode"`

	// tenant ID
	TenantID string `json:"TenantID"`
}

// Validate validates this get tenan user response
func (m *GetTenanUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get tenan user response based on context it is used
func (m *GetTenanUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetTenanUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTenanUserResponse) UnmarshalBinary(b []byte) error {
	var res GetTenanUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
