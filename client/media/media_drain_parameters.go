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

// NewMediaDrainParams creates a new MediaDrainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaDrainParams() *MediaDrainParams {
	return &MediaDrainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaDrainParamsWithTimeout creates a new MediaDrainParams object
// with the ability to set a timeout on a request.
func NewMediaDrainParamsWithTimeout(timeout time.Duration) *MediaDrainParams {
	return &MediaDrainParams{
		timeout: timeout,
	}
}

// NewMediaDrainParamsWithContext creates a new MediaDrainParams object
// with the ability to set a context for a request.
func NewMediaDrainParamsWithContext(ctx context.Context) *MediaDrainParams {
	return &MediaDrainParams{
		Context: ctx,
	}
}

// NewMediaDrainParamsWithHTTPClient creates a new MediaDrainParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaDrainParamsWithHTTPClient(client *http.Client) *MediaDrainParams {
	return &MediaDrainParams{
		HTTPClient: client,
	}
}

/*
MediaDrainParams contains all the parameters to send to the API endpoint

	for the media drain operation.

	Typically these are written to a http.Request.
*/
type MediaDrainParams struct {

	// Authorization.
	Authorization *string

	// Media.
	Media string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media drain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaDrainParams) WithDefaults() *MediaDrainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media drain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaDrainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the media drain params
func (o *MediaDrainParams) WithTimeout(timeout time.Duration) *MediaDrainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media drain params
func (o *MediaDrainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media drain params
func (o *MediaDrainParams) WithContext(ctx context.Context) *MediaDrainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media drain params
func (o *MediaDrainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media drain params
func (o *MediaDrainParams) WithHTTPClient(client *http.Client) *MediaDrainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media drain params
func (o *MediaDrainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the media drain params
func (o *MediaDrainParams) WithAuthorization(authorization *string) *MediaDrainParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the media drain params
func (o *MediaDrainParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithMedia adds the media to the media drain params
func (o *MediaDrainParams) WithMedia(media string) *MediaDrainParams {
	o.SetMedia(media)
	return o
}

// SetMedia adds the media to the media drain params
func (o *MediaDrainParams) SetMedia(media string) {
	o.Media = media
}

// WriteToRequest writes these params to a swagger request
func (o *MediaDrainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param authorization
		if err := r.SetHeaderParam("authorization", *o.Authorization); err != nil {
			return err
		}
	}

	// path param media
	if err := r.SetPathParam("media", o.Media); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
