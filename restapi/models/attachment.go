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

// Attachment attachment
//
// swagger:model Attachment
type Attachment struct {

	// Mount point
	MountPoint string `json:"mountpoint"`

	// node
	// Required: true
	Node string `json:"node"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// Read Only
	ReadOnly *bool `json:"readonly"`

	// Snapshot ID
	// Read Only: true
	SnapshotID string `json:"snapshotid"`

	// Snapshot Name
	SnapshotName string `json:"snapshotname"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// Volume ID
	// Read Only: true
	VolumeID string `json:"volumeid"`

	// Volume name
	// Example: vol1
	VolumeName string `json:"volumename"`
}

// Validate validates this attachment
func (m *Attachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Attachment) validateNode(formats strfmt.Registry) error {

	if err := validate.RequiredString("node", "body", m.Node); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this attachment based on the context it is used
func (m *Attachment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Attachment) contextValidateProgress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "progress", "body", int64(m.Progress)); err != nil {
		return err
	}

	return nil
}

func (m *Attachment) contextValidateSnapshotID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snapshotid", "body", string(m.SnapshotID)); err != nil {
		return err
	}

	return nil
}

func (m *Attachment) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Attachment) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *Attachment) contextValidateVolumeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volumeid", "body", string(m.VolumeID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Attachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attachment) UnmarshalBinary(b []byte) error {
	var res Attachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
