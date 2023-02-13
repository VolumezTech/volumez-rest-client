// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangePasswordRequestLoggedIn change password request logged in
//
// swagger:model ChangePasswordRequestLoggedIn
type ChangePasswordRequestLoggedIn struct {

	// name
	Name string `json:"name"`

	// newpassword
	Newpassword string `json:"newpassword"`

	// oldpassword
	Oldpassword string `json:"oldpassword"`
}

// Validate validates this change password request logged in
func (m *ChangePasswordRequestLoggedIn) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change password request logged in based on context it is used
func (m *ChangePasswordRequestLoggedIn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangePasswordRequestLoggedIn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangePasswordRequestLoggedIn) UnmarshalBinary(b []byte) error {
	var res ChangePasswordRequestLoggedIn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
