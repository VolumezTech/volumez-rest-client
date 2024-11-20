// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new attachments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for attachments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AttachmentCreate(params *AttachmentCreateParams, opts ...ClientOption) (*AttachmentCreateOK, error)

	AttachmentDelete(params *AttachmentDeleteParams, opts ...ClientOption) (*AttachmentDeleteOK, error)

	AttachmentGet(params *AttachmentGetParams, opts ...ClientOption) (*AttachmentGetOK, error)

	AttachmentModify(params *AttachmentModifyParams, opts ...ClientOption) (*AttachmentModifyOK, error)

	AttachmentsList(params *AttachmentsListParams, opts ...ClientOption) (*AttachmentsListOK, error)

	AttachmentsListAll(params *AttachmentsListAllParams, opts ...ClientOption) (*AttachmentsListAllOK, error)

	AttachmentsListForVolume(params *AttachmentsListForVolumeParams, opts ...ClientOption) (*AttachmentsListForVolumeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AttachmentCreate creates a new attachment
*/
func (a *Client) AttachmentCreate(params *AttachmentCreateParams, opts ...ClientOption) (*AttachmentCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachmentCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachmentCreate",
		Method:             "POST",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}/attachments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachmentCreateReader{formats: a.formats},
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
	success, ok := result.(*AttachmentCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachmentCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachmentDelete deletes an attachment
*/
func (a *Client) AttachmentDelete(params *AttachmentDeleteParams, opts ...ClientOption) (*AttachmentDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachmentDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachmentDelete",
		Method:             "DELETE",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}/attachments/{node}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachmentDeleteReader{formats: a.formats},
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
	success, ok := result.(*AttachmentDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachmentDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachmentGet gets the properties of an attachment
*/
func (a *Client) AttachmentGet(params *AttachmentGetParams, opts ...ClientOption) (*AttachmentGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachmentGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachmentGet",
		Method:             "GET",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}/attachments/{node}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachmentGetReader{formats: a.formats},
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
	success, ok := result.(*AttachmentGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachmentGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachmentModify modifies an attachment
*/
func (a *Client) AttachmentModify(params *AttachmentModifyParams, opts ...ClientOption) (*AttachmentModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachmentModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachmentModify",
		Method:             "PATCH",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}/attachments/{node}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachmentModifyReader{formats: a.formats},
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
	success, ok := result.(*AttachmentModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachmentModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachmentsList gets a list of attachments for a given volume and snapshot
*/
func (a *Client) AttachmentsList(params *AttachmentsListParams, opts ...ClientOption) (*AttachmentsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachmentsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachmentsList",
		Method:             "GET",
		PathPattern:        "/volumes/{volume}/snapshots/{snapshot}/attachments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachmentsListReader{formats: a.formats},
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
	success, ok := result.(*AttachmentsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachmentsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachmentsListAll gets a list of all attachments
*/
func (a *Client) AttachmentsListAll(params *AttachmentsListAllParams, opts ...ClientOption) (*AttachmentsListAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachmentsListAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachmentsListAll",
		Method:             "GET",
		PathPattern:        "/attachments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachmentsListAllReader{formats: a.formats},
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
	success, ok := result.(*AttachmentsListAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachmentsListAllDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AttachmentsListForVolume gets a list of attachments for a given volume
*/
func (a *Client) AttachmentsListForVolume(params *AttachmentsListForVolumeParams, opts ...ClientOption) (*AttachmentsListForVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachmentsListForVolumeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AttachmentsListForVolume",
		Method:             "GET",
		PathPattern:        "/volumes/{volume}/attachments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachmentsListForVolumeReader{formats: a.formats},
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
	success, ok := result.(*AttachmentsListForVolumeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AttachmentsListForVolumeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
