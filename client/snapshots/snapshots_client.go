// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new snapshots API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for snapshots API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SnapshotCreate(params *SnapshotCreateParams, opts ...ClientOption) (*SnapshotCreateOK, error)

	SnapshotDelete(params *SnapshotDeleteParams, opts ...ClientOption) (*SnapshotDeleteOK, error)

	SnapshotGet(params *SnapshotGetParams, opts ...ClientOption) (*SnapshotGetOK, error)

	SnapshotModify(params *SnapshotModifyParams, opts ...ClientOption) (*SnapshotModifyOK, error)

	SnapshotRollback(params *SnapshotRollbackParams, opts ...ClientOption) (*SnapshotRollbackOK, error)

	SnapshotsList(params *SnapshotsListParams, opts ...ClientOption) (*SnapshotsListOK, error)

	SnapshotsListAll(params *SnapshotsListAllParams, opts ...ClientOption) (*SnapshotsListAllOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SnapshotCreate creates a new snapshot
*/
func (a *Client) SnapshotCreate(params *SnapshotCreateParams, opts ...ClientOption) (*SnapshotCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SnapshotCreate",
		Method:             "POST",
		PathPattern:        "/volumes/{volume}/snapshots",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotCreateReader{formats: a.formats},
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
	success, ok := result.(*SnapshotCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SnapshotCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SnapshotDelete deletes a snapshot
*/
func (a *Client) SnapshotDelete(params *SnapshotDeleteParams, opts ...ClientOption) (*SnapshotDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SnapshotDelete",
		Method:             "DELETE",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotDeleteReader{formats: a.formats},
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
	success, ok := result.(*SnapshotDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SnapshotDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SnapshotGet gets the properties of a snapshot
*/
func (a *Client) SnapshotGet(params *SnapshotGetParams, opts ...ClientOption) (*SnapshotGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SnapshotGet",
		Method:             "GET",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotGetReader{formats: a.formats},
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
	success, ok := result.(*SnapshotGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SnapshotGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SnapshotModify modifies a snapshot
*/
func (a *Client) SnapshotModify(params *SnapshotModifyParams, opts ...ClientOption) (*SnapshotModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SnapshotModify",
		Method:             "PATCH",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotModifyReader{formats: a.formats},
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
	success, ok := result.(*SnapshotModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SnapshotModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SnapshotRollback rolls back to snapshot
*/
func (a *Client) SnapshotRollback(params *SnapshotRollbackParams, opts ...ClientOption) (*SnapshotRollbackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotRollbackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SnapshotRollback",
		Method:             "PATCH",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}/rollback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotRollbackReader{formats: a.formats},
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
	success, ok := result.(*SnapshotRollbackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SnapshotRollbackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SnapshotsList gets a list of snapshots
*/
func (a *Client) SnapshotsList(params *SnapshotsListParams, opts ...ClientOption) (*SnapshotsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SnapshotsList",
		Method:             "GET",
		PathPattern:        "/volumes/{volume}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotsListReader{formats: a.formats},
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
	success, ok := result.(*SnapshotsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SnapshotsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  SnapshotsListAll gets a list of all snapshots
*/
func (a *Client) SnapshotsListAll(params *SnapshotsListAllParams, opts ...ClientOption) (*SnapshotsListAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSnapshotsListAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SnapshotsListAll",
		Method:             "GET",
		PathPattern:        "/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SnapshotsListAllReader{formats: a.formats},
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
	success, ok := result.(*SnapshotsListAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SnapshotsListAllDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
