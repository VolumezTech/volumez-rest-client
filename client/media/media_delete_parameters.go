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
	"github.com/go-openapi/swag"
)

// NewMediaDeleteParams creates a new MediaDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaDeleteParams() *MediaDeleteParams {
	return &MediaDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaDeleteParamsWithTimeout creates a new MediaDeleteParams object
// with the ability to set a timeout on a request.
func NewMediaDeleteParamsWithTimeout(timeout time.Duration) *MediaDeleteParams {
	return &MediaDeleteParams{
		timeout: timeout,
	}
}

// NewMediaDeleteParamsWithContext creates a new MediaDeleteParams object
// with the ability to set a context for a request.
func NewMediaDeleteParamsWithContext(ctx context.Context) *MediaDeleteParams {
	return &MediaDeleteParams{
		Context: ctx,
	}
}

// NewMediaDeleteParamsWithHTTPClient creates a new MediaDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaDeleteParamsWithHTTPClient(client *http.Client) *MediaDeleteParams {
	return &MediaDeleteParams{
		HTTPClient: client,
	}
}

/*
MediaDeleteParams contains all the parameters to send to the API endpoint

	for the media delete operation.

	Typically these are written to a http.Request.
*/
type MediaDeleteParams struct {

	// Authorization.
	Authorization *string

	// Force.
	Force *bool

	// Media.
	Media string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaDeleteParams) WithDefaults() *MediaDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaDeleteParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := MediaDeleteParams{
		Force: &forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the media delete params
func (o *MediaDeleteParams) WithTimeout(timeout time.Duration) *MediaDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media delete params
func (o *MediaDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media delete params
func (o *MediaDeleteParams) WithContext(ctx context.Context) *MediaDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media delete params
func (o *MediaDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media delete params
func (o *MediaDeleteParams) WithHTTPClient(client *http.Client) *MediaDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media delete params
func (o *MediaDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the media delete params
func (o *MediaDeleteParams) WithAuthorization(authorization *string) *MediaDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the media delete params
func (o *MediaDeleteParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithForce adds the force to the media delete params
func (o *MediaDeleteParams) WithForce(force *bool) *MediaDeleteParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the media delete params
func (o *MediaDeleteParams) SetForce(force *bool) {
	o.Force = force
}

// WithMedia adds the media to the media delete params
func (o *MediaDeleteParams) WithMedia(media string) *MediaDeleteParams {
	o.SetMedia(media)
	return o
}

// SetMedia adds the media to the media delete params
func (o *MediaDeleteParams) SetMedia(media string) {
	o.Media = media
}

// WriteToRequest writes these params to a swagger request
func (o *MediaDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
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
