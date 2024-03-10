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

// Media media
//
// swagger:model Media
type Media struct {

	// bandwidth read
	BandwidthRead int64 `json:"bandwidthread"`

	// bandwidth reserved
	BandwidthReserved int64 `json:"BandwidthReserved"`

	// bandwidth write
	BandwidthWrite int64 `json:"bandwidthwrite"`

	// the capacity group this media belongs to
	// Read Only: true
	CapacityGroup string `json:"capacitygroup"`

	// cloud provider
	CloudProvider string `json:"cloudprovider,omitempty"`

	// the media node FaultDomain
	FaultDomain string `json:"FaultDomain,omitempty"`

	// free bandwidth read
	FreeBandwidthRead int64 `json:"freebandwidthread"`

	// free bandwidth write
	FreeBandwidthWrite int64 `json:"freebandwidthwrite"`

	// free i o p s read
	FreeIOPSRead int64 `json:"freeiopsread"`

	// free i o p s write
	FreeIOPSWrite int64 `json:"freeiopswrite"`

	// free size
	FreeSize int64 `json:"freesize"`

	// i o p s read
	IOPSRead int64 `json:"iopsread"`

	// i o p s write
	IOPSWrite int64 `json:"iopswrite"`

	// latency read
	LatencyRead int64 `json:"latencyread"`

	// latency write
	LatencyWrite int64 `json:"latencywrite"`

	// media ID
	MediaID string `json:"mediaid"`

	// offline time
	OfflineTime DateTime `json:"offlinetime"`

	// the media node PhysicalProximityGroup
	PhysicalProximityGroup string `json:"PhysicalProximityGroup,omitempty"`

	// the media node ResiliencyDomain
	ResiliencyDomain string `json:"ResiliencyDomain,omitempty"`

	// the media node ResourceNamespace
	ResourceNamespace string `json:"ResourceNamespace,omitempty"`

	// s e d
	SED bool `json:"sed"`

	// sector size
	SectorSize int64 `json:"sectorsize"`

	// count of how many volumes are using the media
	// Read Only: true
	VolumesCount int64 `json:"volumescount"`

	// the media node ResourceNamespace
	AccountID string `json:"accountID,omitempty"`

	// assignment
	// Read Only: true
	Assignment string `json:"assignment"`

	// bus
	Bus string `json:"bus"`

	// firmware
	Firmware string `json:"firmware"`

	// location
	Location string `json:"location"`

	// media
	Media string `json:"media"`

	// model
	Model string `json:"model"`

	// node
	Node string `json:"node"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// region
	// Example: us-east-1
	Region string `json:"region"`

	// size
	Size int64 `json:"size"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// zone
	Zone string `json:"zone"`
}

// Validate validates this media
func (m *Media) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOfflineTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Media) validateOfflineTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OfflineTime) { // not required
		return nil
	}

	if err := m.OfflineTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("offlinetime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("offlinetime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this media based on the context it is used
func (m *Media) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacityGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumesCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssignment(ctx, formats); err != nil {
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

func (m *Media) contextValidateCapacityGroup(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "capacitygroup", "body", string(m.CapacityGroup)); err != nil {
		return err
	}

	return nil
}

func (m *Media) contextValidateVolumesCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "volumescount", "body", int64(m.VolumesCount)); err != nil {
		return err
	}

	return nil
}

func (m *Media) contextValidateAssignment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "assignment", "body", string(m.Assignment)); err != nil {
		return err
	}

	return nil
}

func (m *Media) contextValidateProgress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "progress", "body", int64(m.Progress)); err != nil {
		return err
	}

	return nil
}

func (m *Media) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Media) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Media) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Media) UnmarshalBinary(b []byte) error {
	var res Media
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
