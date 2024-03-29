// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewMediaListParams creates a new MediaListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaListParams() *MediaListParams {
	return &MediaListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaListParamsWithTimeout creates a new MediaListParams object
// with the ability to set a timeout on a request.
func NewMediaListParamsWithTimeout(timeout time.Duration) *MediaListParams {
	return &MediaListParams{
		timeout: timeout,
	}
}

// NewMediaListParamsWithContext creates a new MediaListParams object
// with the ability to set a context for a request.
func NewMediaListParamsWithContext(ctx context.Context) *MediaListParams {
	return &MediaListParams{
		Context: ctx,
	}
}

// NewMediaListParamsWithHTTPClient creates a new MediaListParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaListParamsWithHTTPClient(client *http.Client) *MediaListParams {
	return &MediaListParams{
		HTTPClient: client,
	}
}

/*
MediaListParams contains all the parameters to send to the API endpoint

	for the media list operation.

	Typically these are written to a http.Request.
*/
type MediaListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaListParams) WithDefaults() *MediaListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the media list params
func (o *MediaListParams) WithTimeout(timeout time.Duration) *MediaListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media list params
func (o *MediaListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media list params
func (o *MediaListParams) WithContext(ctx context.Context) *MediaListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media list params
func (o *MediaListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media list params
func (o *MediaListParams) WithHTTPClient(client *http.Client) *MediaListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media list params
func (o *MediaListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MediaListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
