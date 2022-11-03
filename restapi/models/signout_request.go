// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SignoutRequest signout request
//
// swagger:model SignoutRequest
type SignoutRequest struct {

	// access token
	// Required: true
	// Min Length: 1
	AccessToken string `json:"AccessToken"`
}

// Validate validates this signout request
func (m *SignoutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignoutRequest) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.RequiredString("AccessToken", "body", m.AccessToken); err != nil {
		return err
	}

	if err := validate.MinLength("AccessToken", "body", m.AccessToken, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signout request based on context it is used
func (m *SignoutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignoutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignoutRequest) UnmarshalBinary(b []byte) error {
	var res SignoutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
