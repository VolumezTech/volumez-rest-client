// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Snapshot snapshot
//
// swagger:model Snapshot
type Snapshot struct {

	// consistency
	// Required: true
	// Enum: [crash application]
	Consistency string `json:"consistency"`

	// consistency group
	ConsistencyGroup bool `json:"consistencygroup"`

	// consistency group name
	ConsistencyGroupName *string `json:"consistencygroupname,omitempty"`

	// number of attachments
	// Read Only: true
	NumberOfAttachments int64 `json:"numberofattachments"`

	// policy
	Policy bool `json:"policy"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// Snapshot name
	// Required: true
	// Min Length: 1
	SnapName string `json:"name"`

	// Snapshot ID
	// Read Only: true
	// Min Length: 1
	SnapshotID string `json:"snapshotid"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// Target connect secret
	// Min Length: 1
	TargetSecret string `json:"targetsecret"`

	// time
	Time DateTime `json:"time"`

	// used
	Used int64 `json:"used"`

	// Volume ID
	// Read Only: true
	// Min Length: 1
	VolumeID string `json:"volumeid"`

	// Volume name
	// Example: vol1
	VolumeName string `json:"volumename"`

	// volume size
	// Read Only: true
	VolumeSize int64 `json:"volumesize"`
}

// Validate validates this snapshot
func (m *Snapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsistency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var snapshotTypeConsistencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["crash","application"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotTypeConsistencyPropEnum = append(snapshotTypeConsistencyPropEnum, v)
	}
}

const (

	// SnapshotConsistencyCrash captures enum value "crash"
	SnapshotConsistencyCrash string = "crash"

	// SnapshotConsistencyApplication captures enum value "application"
	SnapshotConsistencyApplication string = "application"
)

// prop value enum
func (m *Snapshot) validateConsistencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotTypeConsistencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Snapshot) validateConsistency(formats strfmt.Registry) error {

	if err := validate.RequiredString("consistency", "body", m.Consistency); err != nil {
		return err
	}

	// value enum
	if err := m.validateConsistencyEnum("consistency", "body", m.Consistency); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) validateSnapName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.SnapName); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.SnapName, 1); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) validateSnapshotID(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotID) { // not required
		return nil
	}

	if err := validate.MinLength("snapshotid", "body", m.SnapshotID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) validateTargetSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetSecret) { // not required
		return nil
	}

	if err := validate.MinLength("targetsecret", "body", m.TargetSecret, 1); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("time")
		}
		return err
	}

	return nil
}

func (m *Snapshot) validateVolumeID(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeID) { // not required
		return nil
	}

	if err := validate.MinLength("volumeid", "body", m.VolumeID, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapshot based on the context it is used
func (m *Snapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNumberOfAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.contextValidateVolumeSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snapshot) contextValidateNumberOfAttachments(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numberofattachments", "body", int64(m.NumberOfAttachments)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateProgress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "progress", "body", int64(m.Progress)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateSnapshotID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snapshotid", "body", string(m.SnapshotID)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateVolumeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volumeid", "body", string(m.VolumeID)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateVolumeSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volumesize", "body", int64(m.VolumeSize)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Snapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Snapshot) UnmarshalBinary(b []byte) error {
	var res Snapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
