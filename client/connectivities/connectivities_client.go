// Code generated by go-swagger; DO NOT EDIT.

package connectivities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new connectivities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for connectivities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConnectivitiesList(params *ConnectivitiesListParams, opts ...ClientOption) (*ConnectivitiesListOK, error)

	ConnectivityCreate(params *ConnectivityCreateParams, opts ...ClientOption) (*ConnectivityCreateOK, error)

	ConnectivityDelete(params *ConnectivityDeleteParams, opts ...ClientOption) (*ConnectivityDeleteOK, error)

	ConnectivityGet(params *ConnectivityGetParams, opts ...ClientOption) (*ConnectivityGetOK, error)

	ConnectivityModify(params *ConnectivityModifyParams, opts ...ClientOption) (*ConnectivityModifyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ConnectivitiesList gets a list of connectivities
*/
func (a *Client) ConnectivitiesList(params *ConnectivitiesListParams, opts ...ClientOption) (*ConnectivitiesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectivitiesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectivitiesList",
		Method:             "GET",
		PathPattern:        "/connectivities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConnectivitiesListReader{formats: a.formats},
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
	success, ok := result.(*ConnectivitiesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectivitiesListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ConnectivityCreate creates a new connectivity
*/
func (a *Client) ConnectivityCreate(params *ConnectivityCreateParams, opts ...ClientOption) (*ConnectivityCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectivityCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectivityCreate",
		Method:             "POST",
		PathPattern:        "/connectivities",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConnectivityCreateReader{formats: a.formats},
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
	success, ok := result.(*ConnectivityCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectivityCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ConnectivityDelete deletes a connectivity
*/
func (a *Client) ConnectivityDelete(params *ConnectivityDeleteParams, opts ...ClientOption) (*ConnectivityDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectivityDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectivityDelete",
		Method:             "DELETE",
		PathPattern:        "/connectivities/{connectivity}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConnectivityDeleteReader{formats: a.formats},
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
	success, ok := result.(*ConnectivityDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectivityDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ConnectivityGet gets the properties of a connectivity
*/
func (a *Client) ConnectivityGet(params *ConnectivityGetParams, opts ...ClientOption) (*ConnectivityGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectivityGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectivityGet",
		Method:             "GET",
		PathPattern:        "/connectivities/{connectivity}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConnectivityGetReader{formats: a.formats},
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
	success, ok := result.(*ConnectivityGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectivityGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ConnectivityModify modifies a connectivity
*/
func (a *Client) ConnectivityModify(params *ConnectivityModifyParams, opts ...ClientOption) (*ConnectivityModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectivityModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectivityModify",
		Method:             "PATCH",
		PathPattern:        "/connectivities/{connectivity}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConnectivityModifyReader{formats: a.formats},
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
	success, ok := result.(*ConnectivityModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectivityModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
