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

// NewMediaProfileModifyParams creates a new MediaProfileModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaProfileModifyParams() *MediaProfileModifyParams {
	return &MediaProfileModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaProfileModifyParamsWithTimeout creates a new MediaProfileModifyParams object
// with the ability to set a timeout on a request.
func NewMediaProfileModifyParamsWithTimeout(timeout time.Duration) *MediaProfileModifyParams {
	return &MediaProfileModifyParams{
		timeout: timeout,
	}
}

// NewMediaProfileModifyParamsWithContext creates a new MediaProfileModifyParams object
// with the ability to set a context for a request.
func NewMediaProfileModifyParamsWithContext(ctx context.Context) *MediaProfileModifyParams {
	return &MediaProfileModifyParams{
		Context: ctx,
	}
}

// NewMediaProfileModifyParamsWithHTTPClient creates a new MediaProfileModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaProfileModifyParamsWithHTTPClient(client *http.Client) *MediaProfileModifyParams {
	return &MediaProfileModifyParams{
		HTTPClient: client,
	}
}

/*
MediaProfileModifyParams contains all the parameters to send to the API endpoint

	for the media profile modify operation.

	Typically these are written to a http.Request.
*/
type MediaProfileModifyParams struct {

	// Authorization.
	Authorization *string

	/* Body.

	   A Media Profile object
	*/
	Body *models.MediaProfile

	// Media.
	Media string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media profile modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaProfileModifyParams) WithDefaults() *MediaProfileModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media profile modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaProfileModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the media profile modify params
func (o *MediaProfileModifyParams) WithTimeout(timeout time.Duration) *MediaProfileModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media profile modify params
func (o *MediaProfileModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media profile modify params
func (o *MediaProfileModifyParams) WithContext(ctx context.Context) *MediaProfileModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media profile modify params
func (o *MediaProfileModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media profile modify params
func (o *MediaProfileModifyParams) WithHTTPClient(client *http.Client) *MediaProfileModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media profile modify params
func (o *MediaProfileModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the media profile modify params
func (o *MediaProfileModifyParams) WithAuthorization(authorization *string) *MediaProfileModifyParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the media profile modify params
func (o *MediaProfileModifyParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithBody adds the body to the media profile modify params
func (o *MediaProfileModifyParams) WithBody(body *models.MediaProfile) *MediaProfileModifyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the media profile modify params
func (o *MediaProfileModifyParams) SetBody(body *models.MediaProfile) {
	o.Body = body
}

// WithMedia adds the media to the media profile modify params
func (o *MediaProfileModifyParams) WithMedia(media string) *MediaProfileModifyParams {
	o.SetMedia(media)
	return o
}

// SetMedia adds the media to the media profile modify params
func (o *MediaProfileModifyParams) SetMedia(media string) {
	o.Media = media
}

// WriteToRequest writes these params to a swagger request
func (o *MediaProfileModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
