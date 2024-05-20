// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new media API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new media API client with basic auth credentials.
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

// New creates a new media API client with a bearer token for authentication.
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
Client for media API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	MediaAssign(params *MediaAssignParams, opts ...ClientOption) (*MediaAssignOK, error)

	MediaDelete(params *MediaDeleteParams, opts ...ClientOption) (*MediaDeleteOK, error)

	MediaGet(params *MediaGetParams, opts ...ClientOption) (*MediaGetOK, error)

	MediaLedOff(params *MediaLedOffParams, opts ...ClientOption) (*MediaLedOffOK, error)

	MediaLedOn(params *MediaLedOnParams, opts ...ClientOption) (*MediaLedOnOK, error)

	MediaList(params *MediaListParams, opts ...ClientOption) (*MediaListOK, error)

	MediaModify(params *MediaModifyParams, opts ...ClientOption) (*MediaModifyOK, error)

	MediaProfile(params *MediaProfileParams, opts ...ClientOption) (*MediaProfileOK, error)

	MediaProfileModify(params *MediaProfileModifyParams, opts ...ClientOption) (*MediaProfileModifyOK, error)

	MediaUnassign(params *MediaUnassignParams, opts ...ClientOption) (*MediaUnassignOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
MediaAssign assigns media
*/
func (a *Client) MediaAssign(params *MediaAssignParams, opts ...ClientOption) (*MediaAssignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaAssignParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaAssign",
		Method:             "GET",
		PathPattern:        "/media/{media}/assign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaAssignReader{formats: a.formats},
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
	success, ok := result.(*MediaAssignOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaAssignDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaDelete deletes a media
*/
func (a *Client) MediaDelete(params *MediaDeleteParams, opts ...ClientOption) (*MediaDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaDelete",
		Method:             "DELETE",
		PathPattern:        "/media/{media}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaDeleteReader{formats: a.formats},
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
	success, ok := result.(*MediaDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaGet gets the properties of a media
*/
func (a *Client) MediaGet(params *MediaGetParams, opts ...ClientOption) (*MediaGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaGet",
		Method:             "GET",
		PathPattern:        "/media/{media}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaGetReader{formats: a.formats},
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
	success, ok := result.(*MediaGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaLedOff turns the media l e d off
*/
func (a *Client) MediaLedOff(params *MediaLedOffParams, opts ...ClientOption) (*MediaLedOffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaLedOffParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaLedOff",
		Method:             "GET",
		PathPattern:        "/media/{media}/ledoff",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaLedOffReader{formats: a.formats},
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
	success, ok := result.(*MediaLedOffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaLedOffDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaLedOn turns the media l e d on
*/
func (a *Client) MediaLedOn(params *MediaLedOnParams, opts ...ClientOption) (*MediaLedOnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaLedOnParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaLedOn",
		Method:             "GET",
		PathPattern:        "/media/{media}/ledon",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaLedOnReader{formats: a.formats},
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
	success, ok := result.(*MediaLedOnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaLedOnDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaList gets a list of media
*/
func (a *Client) MediaList(params *MediaListParams, opts ...ClientOption) (*MediaListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaList",
		Method:             "GET",
		PathPattern:        "/media",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaListReader{formats: a.formats},
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
	success, ok := result.(*MediaListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaModify modifies a media properties
*/
func (a *Client) MediaModify(params *MediaModifyParams, opts ...ClientOption) (*MediaModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaModify",
		Method:             "PATCH",
		PathPattern:        "/media/{media}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaModifyReader{formats: a.formats},
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
	success, ok := result.(*MediaModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaProfile profiles a media deprecated
*/
func (a *Client) MediaProfile(params *MediaProfileParams, opts ...ClientOption) (*MediaProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaProfile",
		Method:             "GET",
		PathPattern:        "/media/{media}/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaProfileReader{formats: a.formats},
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
	success, ok := result.(*MediaProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaProfileModify modifies a media profile
*/
func (a *Client) MediaProfileModify(params *MediaProfileModifyParams, opts ...ClientOption) (*MediaProfileModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaProfileModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaProfileModify",
		Method:             "PATCH",
		PathPattern:        "/media/{media}/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaProfileModifyReader{formats: a.formats},
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
	success, ok := result.(*MediaProfileModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaProfileModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MediaUnassign unassigns media
*/
func (a *Client) MediaUnassign(params *MediaUnassignParams, opts ...ClientOption) (*MediaUnassignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaUnassignParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MediaUnassign",
		Method:             "GET",
		PathPattern:        "/media/{media}/unassign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaUnassignReader{formats: a.formats},
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
	success, ok := result.(*MediaUnassignOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MediaUnassignDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
