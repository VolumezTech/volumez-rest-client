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

// VolumeGroup volume group
//
// swagger:model VolumeGroup
type VolumeGroup struct {

	// allocated size
	AllocatedSize int64 `json:"allocatedsize,omitempty"`

	// allocation
	Allocation string `json:"allocation,omitempty"`

	// cache media
	CacheMedia []string `json:"cachemedia"`

	// cache media bandwidth read
	CacheMediaBandwidthRead int64 `json:"cachemediabandwidthread,omitempty"`

	// cache media bandwidth write
	CacheMediaBandwidthWrite int64 `json:"cachemediabandwidthwrite,omitempty"`

	// cache media i o p s read
	CacheMediaIOPSRead int64 `json:"cachemediaiopsread,omitempty"`

	// cache media i o p s write
	CacheMediaIOPSWrite int64 `json:"cachemediaiopswrite,omitempty"`

	// cache media size
	CacheMediaSize int64 `json:"cachemediasize,omitempty"`

	// cache raid columns
	CacheRaidColumns int64 `json:"cacheraidcolumns,omitempty"`

	// cache redundancy
	CacheRedundancy int64 `json:"cacheredundancy,omitempty"`

	// cache resiliency
	CacheResiliency string `json:"cacheresiliency,omitempty"`

	// cache size
	CacheSize int64 `json:"cachesize,omitempty"`

	// compression
	Compression bool `json:"compression,omitempty"`

	// deduplication
	Deduplication bool `json:"deduplication,omitempty"`

	// encryption
	Encryption bool `json:"encryption,omitempty"`

	// media
	Media []string `json:"media"`

	// media bandwidth read
	MediaBandwidthRead int64 `json:"mediabandwidthread,omitempty"`

	// media bandwidth write
	MediaBandwidthWrite int64 `json:"mediabandwidthwrite,omitempty"`

	// media i o p s read
	MediaIOPSRead int64 `json:"mediaiopsread,omitempty"`

	// media i o p s write
	MediaIOPSWrite int64 `json:"mediaiopswrite,omitempty"`

	// media size
	MediaSize int64 `json:"mediasize,omitempty"`

	// raid columns
	RaidColumns int64 `json:"raidcolumns,omitempty"`

	// redundancy
	Redundancy int64 `json:"redundancy,omitempty"`

	// resiliency
	Resiliency string `json:"resiliency,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// Target connect secret
	// Min Length: 1
	TargetSecret string `json:"targetsecret"`

	// volume group name
	// Read Only: true
	VolumeGroupName string `json:"Volumegroupname,omitempty"`

	// write cache
	WriteCache bool `json:"writecache,omitempty"`

	// integrity
	Integrity bool `json:"Integrity,omitempty"`
}

// Validate validates this volume group
func (m *VolumeGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroup) validateTargetSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetSecret) { // not required
		return nil
	}

	if err := validate.MinLength("targetsecret", "body", m.TargetSecret, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this volume group based on the context it is used
func (m *VolumeGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumeGroupName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroup) contextValidateVolumeGroupName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "Volumegroupname", "body", string(m.VolumeGroupName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeGroup) UnmarshalBinary(b []byte) error {
	var res VolumeGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
