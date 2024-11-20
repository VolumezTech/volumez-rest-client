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

	"github.com/VolumezTech/volumez-rest-client/restapi/models"
)

// NewMediaAssignParams creates a new MediaAssignParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaAssignParams() *MediaAssignParams {
	return &MediaAssignParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaAssignParamsWithTimeout creates a new MediaAssignParams object
// with the ability to set a timeout on a request.
func NewMediaAssignParamsWithTimeout(timeout time.Duration) *MediaAssignParams {
	return &MediaAssignParams{
		timeout: timeout,
	}
}

// NewMediaAssignParamsWithContext creates a new MediaAssignParams object
// with the ability to set a context for a request.
func NewMediaAssignParamsWithContext(ctx context.Context) *MediaAssignParams {
	return &MediaAssignParams{
		Context: ctx,
	}
}

// NewMediaAssignParamsWithHTTPClient creates a new MediaAssignParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaAssignParamsWithHTTPClient(client *http.Client) *MediaAssignParams {
	return &MediaAssignParams{
		HTTPClient: client,
	}
}

/*
MediaAssignParams contains all the parameters to send to the API endpoint

	for the media assign operation.

	Typically these are written to a http.Request.
*/
type MediaAssignParams struct {

	// Authorization.
	Authorization *string

	// Body.
	Body *models.MediaAssign

	// Media.
	Media string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media assign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaAssignParams) WithDefaults() *MediaAssignParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media assign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaAssignParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the media assign params
func (o *MediaAssignParams) WithTimeout(timeout time.Duration) *MediaAssignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media assign params
func (o *MediaAssignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media assign params
func (o *MediaAssignParams) WithContext(ctx context.Context) *MediaAssignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media assign params
func (o *MediaAssignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media assign params
func (o *MediaAssignParams) WithHTTPClient(client *http.Client) *MediaAssignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media assign params
func (o *MediaAssignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the media assign params
func (o *MediaAssignParams) WithAuthorization(authorization *string) *MediaAssignParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the media assign params
func (o *MediaAssignParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithBody adds the body to the media assign params
func (o *MediaAssignParams) WithBody(body *models.MediaAssign) *MediaAssignParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the media assign params
func (o *MediaAssignParams) SetBody(body *models.MediaAssign) {
	o.Body = body
}

// WithMedia adds the media to the media assign params
func (o *MediaAssignParams) WithMedia(media string) *MediaAssignParams {
	o.SetMedia(media)
	return o
}

// SetMedia adds the media to the media assign params
func (o *MediaAssignParams) SetMedia(media string) {
	o.Media = media
}

// WriteToRequest writes these params to a swagger request
func (o *MediaAssignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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
