// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MediaModify media modify
//
// swagger:model MediaModify
type MediaModify struct {

	// offline time
	OfflineTime *DateTime `json:"offlinetime,omitempty"`
}

// Validate validates this media modify
func (m *MediaModify) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOfflineTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaModify) validateOfflineTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OfflineTime) { // not required
		return nil
	}

	if m.OfflineTime != nil {
		if err := m.OfflineTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("offlinetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("offlinetime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validates this media modify based on context it is used
func (m *MediaModify) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MediaModify) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaModify) UnmarshalBinary(b []byte) error {
	var res MediaModify
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
