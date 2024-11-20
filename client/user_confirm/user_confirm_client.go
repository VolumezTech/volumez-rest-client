// Code generated by go-swagger; DO NOT EDIT.

package user_confirm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user confirm API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user confirm API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserConfirm(params *UserConfirmParams, opts ...ClientOption) (*UserConfirmOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
UserConfirm confirms user signup
*/
func (a *Client) UserConfirm(params *UserConfirmParams, opts ...ClientOption) (*UserConfirmOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserConfirmParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserConfirm",
		Method:             "GET",
		PathPattern:        "/tenant/user/confirmation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserConfirmReader{formats: a.formats},
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
	success, ok := result.(*UserConfirmOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UserConfirmDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
