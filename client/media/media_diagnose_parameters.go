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

// NewMediaDiagnoseParams creates a new MediaDiagnoseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaDiagnoseParams() *MediaDiagnoseParams {
	return &MediaDiagnoseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaDiagnoseParamsWithTimeout creates a new MediaDiagnoseParams object
// with the ability to set a timeout on a request.
func NewMediaDiagnoseParamsWithTimeout(timeout time.Duration) *MediaDiagnoseParams {
	return &MediaDiagnoseParams{
		timeout: timeout,
	}
}

// NewMediaDiagnoseParamsWithContext creates a new MediaDiagnoseParams object
// with the ability to set a context for a request.
func NewMediaDiagnoseParamsWithContext(ctx context.Context) *MediaDiagnoseParams {
	return &MediaDiagnoseParams{
		Context: ctx,
	}
}

// NewMediaDiagnoseParamsWithHTTPClient creates a new MediaDiagnoseParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaDiagnoseParamsWithHTTPClient(client *http.Client) *MediaDiagnoseParams {
	return &MediaDiagnoseParams{
		HTTPClient: client,
	}
}

/*
MediaDiagnoseParams contains all the parameters to send to the API endpoint

	for the media diagnose operation.

	Typically these are written to a http.Request.
*/
type MediaDiagnoseParams struct {

	// Authorization.
	Authorization *string

	// Media.
	Media string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media diagnose params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaDiagnoseParams) WithDefaults() *MediaDiagnoseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media diagnose params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaDiagnoseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the media diagnose params
func (o *MediaDiagnoseParams) WithTimeout(timeout time.Duration) *MediaDiagnoseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media diagnose params
func (o *MediaDiagnoseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media diagnose params
func (o *MediaDiagnoseParams) WithContext(ctx context.Context) *MediaDiagnoseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media diagnose params
func (o *MediaDiagnoseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media diagnose params
func (o *MediaDiagnoseParams) WithHTTPClient(client *http.Client) *MediaDiagnoseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media diagnose params
func (o *MediaDiagnoseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the media diagnose params
func (o *MediaDiagnoseParams) WithAuthorization(authorization *string) *MediaDiagnoseParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the media diagnose params
func (o *MediaDiagnoseParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithMedia adds the media to the media diagnose params
func (o *MediaDiagnoseParams) WithMedia(media string) *MediaDiagnoseParams {
	o.SetMedia(media)
	return o
}

// SetMedia adds the media to the media diagnose params
func (o *MediaDiagnoseParams) SetMedia(media string) {
	o.Media = media
}

// WriteToRequest writes these params to a swagger request
func (o *MediaDiagnoseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
