// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Signout(params *SignoutParams, opts ...ClientOption) (*SignoutOK, error)

	TenantUserCreate(params *TenantUserCreateParams, opts ...ClientOption) (*TenantUserCreateOK, error)

	UserCreate(params *UserCreateParams, opts ...ClientOption) (*UserCreateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Signout signs out user from all devices
*/
func (a *Client) Signout(params *SignoutParams, opts ...ClientOption) (*SignoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSignoutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Signout",
		Method:             "POST",
		PathPattern:        "/signout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SignoutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SignoutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SignoutDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TenantUserCreate creates tenant s additional user
*/
func (a *Client) TenantUserCreate(params *TenantUserCreateParams, opts ...ClientOption) (*TenantUserCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenantUserCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TenantUserCreate",
		Method:             "POST",
		PathPattern:        "/tenant/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenantUserCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TenantUserCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TenantUserCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UserCreate creates tenant s first user
*/
func (a *Client) UserCreate(params *UserCreateParams, opts ...ClientOption) (*UserCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserCreate",
		Method:             "POST",
		PathPattern:        "/signup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
