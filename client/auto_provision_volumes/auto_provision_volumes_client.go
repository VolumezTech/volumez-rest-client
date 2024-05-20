// Code generated by go-swagger; DO NOT EDIT.

package auto_provision_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new auto provision volumes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new auto provision volumes API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new auto provision volumes API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for auto provision volumes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AutoProvisionVolumes(params *AutoProvisionVolumesParams, opts ...ClientOption) (*AutoProvisionVolumesOK, *AutoProvisionVolumesAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AutoProvisionVolumes creates a new auto provisioned volume
*/
func (a *Client) AutoProvisionVolumes(params *AutoProvisionVolumesParams, opts ...ClientOption) (*AutoProvisionVolumesOK, *AutoProvisionVolumesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAutoProvisionVolumesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AutoProvisionVolumes",
		Method:             "POST",
		PathPattern:        "/autoprovisionvolumes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AutoProvisionVolumesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *AutoProvisionVolumesOK:
		return value, nil, nil
	case *AutoProvisionVolumesAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AutoProvisionVolumesDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
