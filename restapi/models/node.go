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

// Node node
//
// swagger:model Node
type Node struct {

	// cloud provider
	CloudProvider string `json:"cloudprovider"`

	// cluster
	Cluster string `json:"nodecluster,omitempty"`

	// connector version
	// Read Only: true
	ConnectorVersion string `json:"connectorversion"`

	// control address
	ControlAddress string `json:"controladdress"`

	// o s
	// Example: rhel
	// Required: true
	OS string `json:"os"`

	// offline time
	OfflineTime DateTime `json:"offlinetime"`

	// credential
	Credential string `json:"credential"`

	// label
	Label *string `json:"label,omitempty"`

	// name
	// Example: rv1
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// progress
	// Read Only: true
	Progress int64 `json:"progress"`

	// region
	// Example: us-east-1
	Region string `json:"region"`

	// state
	// Read Only: true
	State string `json:"state"`

	// status
	// Read Only: true
	Status string `json:"status"`

	// zone
	// Example: z1
	Zone string `json:"zone"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateOS(formats strfmt.Registry) error {

	if err := validate.RequiredString("os", "body", m.OS); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateOfflineTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OfflineTime) { // not required
		return nil
	}

	if err := m.OfflineTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("offlinetime")
		}
		return err
	}

	return nil
}

func (m *Node) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this node based on the context it is used
func (m *Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectorVersion(ctx, formats); err != nil {
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

func (m *Node) contextValidateConnectorVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connectorversion", "body", string(m.ConnectorVersion)); err != nil {
		return err
	}

	return nil
}

func (m *Node) contextValidateProgress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "progress", "body", int64(m.Progress)); err != nil {
		return err
	}

	return nil
}

func (m *Node) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Node) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
