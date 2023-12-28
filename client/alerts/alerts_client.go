// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AlertAcknowledge(params *AlertAcknowledgeParams, opts ...ClientOption) (*AlertAcknowledgeOK, error)

	AlertsList(params *AlertsListParams, opts ...ClientOption) (*AlertsListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AlertAcknowledge acknowledges a pending alert
*/
func (a *Client) AlertAcknowledge(params *AlertAcknowledgeParams, opts ...ClientOption) (*AlertAcknowledgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertAcknowledgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertAcknowledge",
		Method:             "POST",
		PathPattern:        "/alerts/{alert}/acknowledge",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AlertAcknowledgeReader{formats: a.formats},
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
	success, ok := result.(*AlertAcknowledgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AlertAcknowledgeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AlertsList gets a list of alerts
*/
func (a *Client) AlertsList(params *AlertsListParams, opts ...ClientOption) (*AlertsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertsList",
		Method:             "GET",
		PathPattern:        "/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AlertsListReader{formats: a.formats},
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
	success, ok := result.(*AlertsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AlertsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
