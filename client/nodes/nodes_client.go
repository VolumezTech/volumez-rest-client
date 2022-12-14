// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nodes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nodes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	NodeCleanUp(params *NodeCleanUpParams, opts ...ClientOption) (*NodeCleanUpOK, error)

	NodeDelete(params *NodeDeleteParams, opts ...ClientOption) (*NodeDeleteOK, error)

	NodeGet(params *NodeGetParams, opts ...ClientOption) (*NodeGetOK, error)

	NodesList(params *NodesListParams, opts ...ClientOption) (*NodesListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  NodeCleanUp performings node cleanup
*/
func (a *Client) NodeCleanUp(params *NodeCleanUpParams, opts ...ClientOption) (*NodeCleanUpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeCleanUpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeCleanUp",
		Method:             "POST",
		PathPattern:        "/nodes/cleanup/{node}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeCleanUpReader{formats: a.formats},
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
	success, ok := result.(*NodeCleanUpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeCleanUpDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NodeDelete deletes a node
*/
func (a *Client) NodeDelete(params *NodeDeleteParams, opts ...ClientOption) (*NodeDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeDelete",
		Method:             "DELETE",
		PathPattern:        "/nodes/{node}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeDeleteReader{formats: a.formats},
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
	success, ok := result.(*NodeDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NodeGet gets the properties of a node
*/
func (a *Client) NodeGet(params *NodeGetParams, opts ...ClientOption) (*NodeGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeGet",
		Method:             "GET",
		PathPattern:        "/nodes/{node}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeGetReader{formats: a.formats},
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
	success, ok := result.(*NodeGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NodesList gets a list of nodes
*/
func (a *Client) NodesList(params *NodesListParams, opts ...ClientOption) (*NodesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodesList",
		Method:             "GET",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodesListReader{formats: a.formats},
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
	success, ok := result.(*NodesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodesListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
