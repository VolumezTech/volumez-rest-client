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

// AutoProvisionVolume auto provision volume
//
// swagger:model AutoProvisionVolume
type AutoProvisionVolume struct {

	// availability zones
	// Required: true
	AvailabilityZones []string `json:"availabilityZones"`

	// cloud provider name
	// Example: Aws
	// Required: true
	// Min Length: 1
	CloudProvider string `json:"cloudProvider"`

	// os image id
	// Min Length: 1
	ImageID string `json:"imageId"`

	// os type
	// Required: true
	// Enum: [Linux Rhel Ubuntu]
	OsType string `json:"osType"`

	// Region to create teh volume in
	// Example: us-east-1
	// Required: true
	// Min Length: 1
	Region string `json:"region"`

	// ssh key name
	// Min Length: 1
	SSHKeyName string `json:"sshKeyName"`

	// subnets
	// Required: true
	Subnets []string `json:"subnets"`

	// infra plan
	InfraPlan *AutoProvisionInfraPlan `json:"infraPlan,omitempty"`

	// the customers provider account id
	// Required: true
	// Min Length: 1
	ProviderAccountID string `json:"accountId"`

	// volume
	// Required: true
	Volume *Volume `json:"volume"`
}

// Validate validates this auto provision volume
func (m *AutoProvisionVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHKeyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfraPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoProvisionVolume) validateAvailabilityZones(formats strfmt.Registry) error {

	if err := validate.Required("availabilityZones", "body", m.AvailabilityZones); err != nil {
		return err
	}

	return nil
}

func (m *AutoProvisionVolume) validateCloudProvider(formats strfmt.Registry) error {

	if err := validate.RequiredString("cloudProvider", "body", m.CloudProvider); err != nil {
		return err
	}

	if err := validate.MinLength("cloudProvider", "body", m.CloudProvider, 1); err != nil {
		return err
	}

	return nil
}

func (m *AutoProvisionVolume) validateImageID(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageID) { // not required
		return nil
	}

	if err := validate.MinLength("imageId", "body", m.ImageID, 1); err != nil {
		return err
	}

	return nil
}

var autoProvisionVolumeTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Linux","Rhel","Ubuntu"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoProvisionVolumeTypeOsTypePropEnum = append(autoProvisionVolumeTypeOsTypePropEnum, v)
	}
}

const (

	// AutoProvisionVolumeOsTypeLinux captures enum value "Linux"
	AutoProvisionVolumeOsTypeLinux string = "Linux"

	// AutoProvisionVolumeOsTypeRhel captures enum value "Rhel"
	AutoProvisionVolumeOsTypeRhel string = "Rhel"

	// AutoProvisionVolumeOsTypeUbuntu captures enum value "Ubuntu"
	AutoProvisionVolumeOsTypeUbuntu string = "Ubuntu"
)

// prop value enum
func (m *AutoProvisionVolume) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autoProvisionVolumeTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutoProvisionVolume) validateOsType(formats strfmt.Registry) error {

	if err := validate.RequiredString("osType", "body", m.OsType); err != nil {
		return err
	}

	// value enum
	if err := m.validateOsTypeEnum("osType", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *AutoProvisionVolume) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", m.Region); err != nil {
		return err
	}

	if err := validate.MinLength("region", "body", m.Region, 1); err != nil {
		return err
	}

	return nil
}

func (m *AutoProvisionVolume) validateSSHKeyName(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHKeyName) { // not required
		return nil
	}

	if err := validate.MinLength("sshKeyName", "body", m.SSHKeyName, 1); err != nil {
		return err
	}

	return nil
}

func (m *AutoProvisionVolume) validateSubnets(formats strfmt.Registry) error {

	if err := validate.Required("subnets", "body", m.Subnets); err != nil {
		return err
	}

	return nil
}

func (m *AutoProvisionVolume) validateInfraPlan(formats strfmt.Registry) error {
	if swag.IsZero(m.InfraPlan) { // not required
		return nil
	}

	if m.InfraPlan != nil {
		if err := m.InfraPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infraPlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infraPlan")
			}
			return err
		}
	}

	return nil
}

func (m *AutoProvisionVolume) validateProviderAccountID(formats strfmt.Registry) error {

	if err := validate.RequiredString("accountId", "body", m.ProviderAccountID); err != nil {
		return err
	}

	if err := validate.MinLength("accountId", "body", m.ProviderAccountID, 1); err != nil {
		return err
	}

	return nil
}

func (m *AutoProvisionVolume) validateVolume(formats strfmt.Registry) error {

	if err := validate.Required("volume", "body", m.Volume); err != nil {
		return err
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auto provision volume based on the context it is used
func (m *AutoProvisionVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfraPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoProvisionVolume) contextValidateInfraPlan(ctx context.Context, formats strfmt.Registry) error {

	if m.InfraPlan != nil {

		if swag.IsZero(m.InfraPlan) { // not required
			return nil
		}

		if err := m.InfraPlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infraPlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infraPlan")
			}
			return err
		}
	}

	return nil
}

func (m *AutoProvisionVolume) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {

		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoProvisionVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoProvisionVolume) UnmarshalBinary(b []byte) error {
	var res AutoProvisionVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}