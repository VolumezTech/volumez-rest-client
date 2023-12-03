// Code generated by go-swagger; DO NOT EDIT.

package tenant_refresh_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenant refresh token API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenant refresh token API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TenantRefreshToken(params *TenantRefreshTokenParams, opts ...ClientOption) (*TenantRefreshTokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
TenantRefreshToken gets the tenant s refresh token
*/
func (a *Client) TenantRefreshToken(params *TenantRefreshTokenParams, opts ...ClientOption) (*TenantRefreshTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenantRefreshTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TenantRefreshToken",
		Method:             "POST",
		PathPattern:        "/tenant/refreshtoken",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenantRefreshTokenReader{formats: a.formats},
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
	success, ok := result.(*TenantRefreshTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TenantRefreshTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
