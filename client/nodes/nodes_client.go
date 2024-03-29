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

	NodeCollectLogs(params *NodeCollectLogsParams, opts ...ClientOption) (*NodeCollectLogsOK, error)

	NodeDelete(params *NodeDeleteParams, opts ...ClientOption) (*NodeDeleteOK, error)

	NodeGet(params *NodeGetParams, opts ...ClientOption) (*NodeGetOK, error)

	NodeRepair(params *NodeRepairParams, opts ...ClientOption) (*NodeRepairOK, error)

	NodeUpgrade(params *NodeUpgradeParams, opts ...ClientOption) (*NodeUpgradeOK, error)

	NodeUpgradeForSupport(params *NodeUpgradeForSupportParams, opts ...ClientOption) (*NodeUpgradeForSupportOK, error)

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
NodeCollectLogs nodes collect logs
*/
func (a *Client) NodeCollectLogs(params *NodeCollectLogsParams, opts ...ClientOption) (*NodeCollectLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeCollectLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeCollectLogs",
		Method:             "POST",
		PathPattern:        "/nodes/logs/{node}/{tenant}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeCollectLogsReader{formats: a.formats},
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
	success, ok := result.(*NodeCollectLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeCollectLogsDefault)
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
NodeRepair executes commands on node for repair
*/
func (a *Client) NodeRepair(params *NodeRepairParams, opts ...ClientOption) (*NodeRepairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeRepairParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeRepair",
		Method:             "POST",
		PathPattern:        "/nodes/repair/{node}/{tenant}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeRepairReader{formats: a.formats},
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
	success, ok := result.(*NodeRepairOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeRepairDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
NodeUpgrade performings node version upgrade
*/
func (a *Client) NodeUpgrade(params *NodeUpgradeParams, opts ...ClientOption) (*NodeUpgradeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeUpgradeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeUpgrade",
		Method:             "POST",
		PathPattern:        "/nodes/upgrade/{node}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeUpgradeReader{formats: a.formats},
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
	success, ok := result.(*NodeUpgradeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeUpgradeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
NodeUpgradeForSupport performings node version upgrade
*/
func (a *Client) NodeUpgradeForSupport(params *NodeUpgradeForSupportParams, opts ...ClientOption) (*NodeUpgradeForSupportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeUpgradeForSupportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "NodeUpgradeForSupport",
		Method:             "POST",
		PathPattern:        "/nodes/upgrade/{node}/tenant/{tenant}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NodeUpgradeForSupportReader{formats: a.formats},
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
	success, ok := result.(*NodeUpgradeForSupportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NodeUpgradeForSupportDefault)
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
