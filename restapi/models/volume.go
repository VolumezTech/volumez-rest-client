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

// Volume volume
//
// swagger:model Volume
type Volume struct {

	// Capacity group from which the volume is allocated
	// Min Length: 1
	CapacityGroup *string `json:"capacitygroup,omitempty"`

	// consistency group
	ConsistencyGroup string `json:"consistencygroup"`

	// content snapshot
	ContentSnapshot string `json:"contentsnapshot"`

	// content volume
	ContentVolume string `json:"contentvolume"`

	// Upper limit size, GiB
	// Example: 100
	// Read Only: true
	MaxSize int64 `json:"maxsize"`

	// Volume name
	// Example: vol1
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// replication node
	ReplicationNode string `json:"replicationnode"`

	// Replication Volume ID
	// Read Only: true
	// Min Length: 1
	ReplicationVolumeGroupID string `json:"replicationvolumegroupid"`

	// Replication Volume Group name
	// Example: vg_1
	// Min Length: 1
	ReplicationVolumeGroupName string `json:"replicationvolumegroupname"`

	// Throttling Scheme for the volume
	// Min Length: 1
	ThrottlingScheme *string `json:"throttlingscheme,omitempty"`

	// Volume ID
	// Read Only: true
	// Min Length: 1
	VolumeGroupID string `json:"volumegroupid"`

	// Volume Group name
	// Example: vg_1
	// Min Length: 1
	VolumeGroupName string `json:"volumegroupname"`

	// Volume ID
	// Read Only: true
	// Min Length: 1
	VolumeID string `json:"volumeid"`

	// Volume Recovery Job ID
	// Read Only: true
	// Min Length: 1
	VolumeRecoveryJob string `json:"volumerecoveryjob"`

	// node
	// Read Only: true
	Node string `json:"node,omitempty"`

	// policy
	// Required: true
	// Min Length: 1
	Policy string `json:"policy"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// Size, GiB
	// Example: 10
	// Required: true
	// Maximum: 65536
	// Minimum: 1
	Size int64 `json:"size"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// type
	// Required: true
	// Enum: [file block]
	Type string `json:"type"`

	// zone
	// Example: eu-west-2
	// Min Length: 1
	Zone string `json:"zone"`

	// zonereplica
	Zonereplica *string `json:"zonereplica"`
}

// Validate validates this volume
func (m *Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicationVolumeGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicationVolumeGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottlingScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeRecoveryJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) validateCapacityGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.CapacityGroup) { // not required
		return nil
	}

	if err := validate.MinLength("capacitygroup", "body", *m.CapacityGroup, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateReplicationVolumeGroupID(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationVolumeGroupID) { // not required
		return nil
	}

	if err := validate.MinLength("replicationvolumegroupid", "body", m.ReplicationVolumeGroupID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateReplicationVolumeGroupName(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationVolumeGroupName) { // not required
		return nil
	}

	if err := validate.MinLength("replicationvolumegroupname", "body", m.ReplicationVolumeGroupName, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateThrottlingScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.ThrottlingScheme) { // not required
		return nil
	}

	if err := validate.MinLength("throttlingscheme", "body", *m.ThrottlingScheme, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateVolumeGroupID(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeGroupID) { // not required
		return nil
	}

	if err := validate.MinLength("volumegroupid", "body", m.VolumeGroupID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateVolumeGroupName(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeGroupName) { // not required
		return nil
	}

	if err := validate.MinLength("volumegroupname", "body", m.VolumeGroupName, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateVolumeID(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeID) { // not required
		return nil
	}

	if err := validate.MinLength("volumeid", "body", m.VolumeID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateVolumeRecoveryJob(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeRecoveryJob) { // not required
		return nil
	}

	if err := validate.MinLength("volumerecoveryjob", "body", m.VolumeRecoveryJob, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validatePolicy(formats strfmt.Registry) error {

	if err := validate.RequiredString("policy", "body", m.Policy); err != nil {
		return err
	}

	if err := validate.MinLength("policy", "body", m.Policy, 1); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", int64(m.Size)); err != nil {
		return err
	}

	if err := validate.MinimumInt("size", "body", m.Size, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("size", "body", m.Size, 65536, false); err != nil {
		return err
	}

	return nil
}

var volumeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["file","block"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeTypeTypePropEnum = append(volumeTypeTypePropEnum, v)
	}
}

const (

	// VolumeTypeFile captures enum value "file"
	VolumeTypeFile string = "file"

	// VolumeTypeBlock captures enum value "block"
	VolumeTypeBlock string = "block"
)

// prop value enum
func (m *Volume) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, volumeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Volume) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateZone(formats strfmt.Registry) error {
	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if err := validate.MinLength("zone", "body", m.Zone, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this volume based on the context it is used
func (m *Volume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaxSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplicationVolumeGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeRecoveryJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) contextValidateMaxSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "maxsize", "body", int64(m.MaxSize)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateReplicationVolumeGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "replicationvolumegroupid", "body", string(m.ReplicationVolumeGroupID)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateVolumeGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volumegroupid", "body", string(m.VolumeGroupID)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateVolumeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volumeid", "body", string(m.VolumeID)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateVolumeRecoveryJob(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volumerecoveryjob", "body", string(m.VolumeRecoveryJob)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "node", "body", string(m.Node)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateProgress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "progress", "body", int64(m.Progress)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Volume) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Volume) UnmarshalBinary(b []byte) error {
	var res Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
