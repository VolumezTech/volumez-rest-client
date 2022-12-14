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

// SignInResponse sign in response
//
// swagger:model SignInResponse
type SignInResponse struct {

	// access token
	// Required: true
	AccessToken string `json:"AccessToken"`

	// expires in
	// Required: true
	ExpiresIn int64 `json:"ExpiresIn"`

	// Id token
	// Required: true
	IDToken string `json:"IdToken"`

	// refresh token
	// Required: true
	RefreshToken string `json:"RefreshToken"`

	// token type
	// Required: true
	TokenType string `json:"TokenType"`
}

// Validate validates this sign in response
func (m *SignInResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignInResponse) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.RequiredString("AccessToken", "body", m.AccessToken); err != nil {
		return err
	}

	return nil
}

func (m *SignInResponse) validateExpiresIn(formats strfmt.Registry) error {

	if err := validate.Required("ExpiresIn", "body", int64(m.ExpiresIn)); err != nil {
		return err
	}

	return nil
}

func (m *SignInResponse) validateIDToken(formats strfmt.Registry) error {

	if err := validate.RequiredString("IdToken", "body", m.IDToken); err != nil {
		return err
	}

	return nil
}

func (m *SignInResponse) validateRefreshToken(formats strfmt.Registry) error {

	if err := validate.RequiredString("RefreshToken", "body", m.RefreshToken); err != nil {
		return err
	}

	return nil
}

func (m *SignInResponse) validateTokenType(formats strfmt.Registry) error {

	if err := validate.RequiredString("TokenType", "body", m.TokenType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sign in response based on context it is used
func (m *SignInResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignInResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignInResponse) UnmarshalBinary(b []byte) error {
	var res SignInResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
